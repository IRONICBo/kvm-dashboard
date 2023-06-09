package services

import (
	"encoding/json"
	"fmt"
	"kvm-dashboard/consts"
	"kvm-dashboard/model"
	"kvm-dashboard/router/ws"
	"kvm-dashboard/utils"
	"kvm-dashboard/vm/agent"
	"kvm-dashboard/vm/data"
	"strconv"
	"time"
)

// get latest data from influxdb
func (svc *Service) GetLatestSimpleData(uuid string, fields []string) []*model.Metric {
	if len(fields) == 0 {
		return nil
	}

	res := make([]*model.Metric, 0)
	for _, field := range fields {
		res = append(res, svc.Dao.ReadSimpleData(uuid, field, consts.PERIOD_1M, ""))
	}

	return res
}

func (svc *Service) GetSimpleData(uuid, period, agg string, fields []string) []*model.Metric {
	if len(fields) == 0 {
		return nil
	}

	res := make([]*model.Metric, 0)
	for _, field := range fields {
		res = append(res, svc.Dao.ReadVMData(uuid, field, period, agg))
	}

	return res
}

func (svc *Service) StartHostSimpleReport(uuid string) error {
	username := svc.GetMachineInfo(uuid).Username
	password := svc.GetMachineInfo(uuid).Password
	sshPort := svc.GetMachineInfo(uuid).SshPort
	port, _ := strconv.Atoi(sshPort)
	ip := svc.GetMachineInfo(uuid).Ip

	simpleAgent, err := agent.NewSimpleAgent(
		&agent.AgentInfo{
			UUID:     uuid,
			User:     username,
			Password: password,
			Ip:       ip,
			Port:     uint(port),
		},
	)

	if err != nil {
		utils.Log.Error("Can not create simple agent", err)
		return err
	}

	// 1. Start collect goroutine
	simpleAgent.Start(uuid)

	// 2. Start collect goroutine
	go svc.saveAndReportSimpleData(simpleAgent, uuid)

	return nil
}

// Start agent and report Simple data
func (svc *Service) StartSimpleReport(uuid string) error {
	// get ip address
	vminfo := svc.GetVMInfo(uuid)
	username := svc.GetMachineInfo(uuid).Username
	password := svc.GetMachineInfo(uuid).Password
	sshPort := svc.GetMachineInfo(uuid).SshPort
	port, _ := strconv.Atoi(sshPort)

	simpleAgent, err := agent.NewSimpleAgent(
		&agent.AgentInfo{
			UUID:     uuid,
			User:     username,
			Password: password,
			Ip:       vminfo.IpAddress,
			Port:     uint(port),
		},
	)

	if err != nil {
		utils.Log.Error("Can not create simple agent", err)
		return err
	}

	// 1. Start collect goroutine
	simpleAgent.Start(uuid)

	// 2. Start collect goroutine
	go svc.saveAndReportSimpleData(simpleAgent, uuid)

	return nil
}

func (svc *Service) saveAndReportSimpleData(agent *agent.SimpleAgent, uuid string) {
	for {
		select {
		case <-time.After(60 * time.Second):
			utils.Log.Info("Time out")
			return
		case <-agent.Ctx.Done():
			utils.Log.Info(fmt.Sprintf("Stop saveAndReportSimpleData goroutine: %#v", uuid))
			return
		case data := <-agent.SimpleData:
			// report to ws
			simple := model.NewSingleMetric("simple", "", data)
			simpleJsonByte, err := json.Marshal(simple)
			if err != nil {
				utils.Log.Error("Can not marshal simple data", err)
			}
			simpleJson := string(simpleJsonByte)
			err = ws.SimpleWSServer.WriteData(uuid, simpleJson)
			if err != nil {
				utils.Log.Debug("Can not find/write simple data to ws", err)
			}

			// alert data
			go svc.alertSimpleData(data, uuid)

			// write to influxdb
			err = svc.Dao.WriteSimpleData(data, map[string]string{
				"uuid": uuid,
			})
			if err != nil {
				utils.Log.Error("Can not write vm data", err)
			}
		}
	}
}

func (svc *Service) alertSimpleData(simpleData *data.SimpleData, uuid string) error {
	if consts.VM_THRESHOLD[uuid] == nil {
		svc.GetMachineThreshold(uuid)
	}

	cpuUsageThreshold := consts.VM_THRESHOLD[uuid].CpuUsageThreshold
	if int(simpleData.CPUUsage) > cpuUsageThreshold {
		alertData := map[string]interface{}{
			"cpu_usage": simpleData.CPUUsage,
		}

		err := svc.Dao.WriteAlertData(alertData, map[string]string{
			"uuid":      uuid,
			"timestamp": time.Now().Local().String(), // used for dumplicate data
		})
		if err != nil {
			utils.Log.Error("Can not write vm data", err)
			return err
		}
	}

	memUsageThreshold := consts.VM_THRESHOLD[uuid].MemUsageThreshold
	if int(simpleData.MemUsage) > memUsageThreshold {
		alertData := map[string]interface{}{
			"mem_usage": simpleData.MemUsage,
		}

		err := svc.Dao.WriteAlertData(alertData, map[string]string{
			"uuid":      uuid,
			"timestamp": time.Now().Local().String(),
		})
		if err != nil {
			utils.Log.Error("Can not write vm data", err)
			return err
		}
	}

	diskUsageThreshold := consts.VM_THRESHOLD[uuid].DiskUsageThreshold
	if int(simpleData.DiskUsage) > diskUsageThreshold {
		alertData := map[string]interface{}{
			"disk_usage": simpleData.DiskUsage,
		}

		err := svc.Dao.WriteAlertData(alertData, map[string]string{
			"uuid":      uuid,
			"timestamp": time.Now().Local().String(),
		})
		if err != nil {
			utils.Log.Error("Can not write vm data", err)
			return err
		}
	}

	return nil
}

// todo: delete url param
func (svc *Service) StopSimpleReport(uuid string) error {
	err := agent.StopSimpleAgent(uuid)
	if err != nil {
		utils.Log.Error("Can not stop agent", err)
		return err
	}

	return nil
}
