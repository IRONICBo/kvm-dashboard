package services

import (
	"encoding/json"
	"fmt"
	"kvm-dashboard/consts"
	"kvm-dashboard/router/ws"
	"kvm-dashboard/utils"
	"kvm-dashboard/vm/agent"
	"time"
)

func (svc *Service) GetProcessInfo(uuid string) error {
	// get ip address
	vminfo := svc.GetVMInfo(uuid)

	// get process agent
	processAgent, err := agent.NewProcessAgent(
		&agent.AgentInfo{
			UUID:     uuid,
			User:     consts.USERNAME,
			Password: consts.PASSWORD,
			Ip:       vminfo.IpAddress,
			Port:     consts.PORT,
		},
	)
	if err != nil {
		utils.Log.Error("Can not create process agent", err)
		return err
	}

	// 1. start process agent
	processAgent.Start(uuid)

	// get process info
	go svc.reportProcessData(processAgent, uuid)

	return nil
}

func (svc *Service) reportProcessData(agent *agent.ProcessAgent, uuid string) {
	for {
		select {
		case <-time.After(60 * time.Second):
			utils.Log.Info("Time out")
			return
		case <-agent.Ctx.Done():
			utils.Log.Info(fmt.Sprintf("Stop saveAndReportVMData goroutine: %#v", uuid))
			return
		case data := <-agent.ProcessData:
			// report to ws
			processJsonByte, err := json.Marshal(data)
			if err != nil {
				utils.Log.Error("Can not marshal process data", err)
			}
			processJson := string(processJsonByte)

			// write to ws
			ws.ProcessWSServer.WriteData(uuid, processJson)
		}
	}

}

func (svc *Service) StopProcessInfo(uuid string) error {
	err := agent.StopProcessAgent(uuid)
	if err != nil {
		utils.Log.Error("Can not stop agent", err)
		return err
	}

	return nil
}
