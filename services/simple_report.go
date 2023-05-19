package services

import (
	"encoding/json"
	"fmt"
	"kvm-dashboard/consts"
	"kvm-dashboard/model"
	"kvm-dashboard/router/ws"
	"kvm-dashboard/utils"
	"kvm-dashboard/vm/agent"
	"time"
)

// get latest data from influxdb
func (svc *Service) GetLatestSimpleData(uuid string, fields []string) []*model.Metric {
	if len(fields) == 0 {
		return nil
	}

	res := make([]*model.Metric, 0)
	for _, field := range fields {
		res = append(res, svc.Dao.ReadVMData(uuid, field, consts.PERIOD_1M, ""))
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

// Start agent and report Simple data
func (svc *Service) StartSimpleReport(uuid string) error {
	// get ip address
	vminfo := svc.GetVMInfo(uuid)

	simpleAgent, err := agent.NewSimpleAgent(
		&agent.AgentInfo{
			UUID:     uuid,
			User:     consts.USERNAME,
			Password: consts.PASSWORD,
			Ip:       vminfo.IpAddress,
			Port:     consts.PORT,
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
				utils.Log.Error("Can not write simple data to ws", err)
			}

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

// todo: delete url param
func (svc *Service) StopSimpleReport(uuid string) error {
	err := agent.StopSimpleAgent(uuid)
	if err != nil {
		utils.Log.Error("Can not stop agent", err)
		return err
	}

	return nil
}