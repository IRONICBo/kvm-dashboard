package services

import (
	"fmt"
	"kvm-dashboard/consts"
	"kvm-dashboard/model"
	"kvm-dashboard/utils"
	"kvm-dashboard/vm/agent"
	"kvm-dashboard/vm/data"
	"sort"
	"time"

	"libvirt.org/libvirt-go"
)

// get latest data from influxdb
func (svc *Service) GetLatestVMData(uuid string, fields []string) []*model.Metric {
	if len(fields) == 0 {
		return nil
	}

	res := make([]*model.Metric, 0)
	for _, field := range fields {
		res = append(res, svc.Dao.ReadVMData(uuid, field, consts.PERIOD_1M, ""))
	}

	return res
}

func (svc *Service) GetVMData(uuid, period, agg string, fields []string) []*model.Metric {
	if len(fields) == 0 {
		return nil
	}

	res := make([]*model.Metric, 0)
	for _, field := range fields {
		res = append(res, svc.Dao.ReadVMData(uuid, field, period, agg))
	}

	return res
}

// Start agent and report data
func (svc *Service) StartReport(url, uuid string) error {
	libvirtAgent, err := agent.NewLibvirtAgent(url)
	if err != nil {
		utils.Log.Error("Can not create libvirt agent", err)
		return err
	}

	// 1. Start collect goroutine
	libvirtAgent.Start(uuid)

	// 2. Start collect goroutine
	go svc.saveAndReportVMData(libvirtAgent, uuid)

	return nil
}

func (svc *Service) saveAndReportVMData(agent *agent.LibvirtAgent, uuid string) {
	for {
		select {
		case <-time.After(60 * time.Second):
			utils.Log.Info("Time out")
			return
		case <-agent.Ctx.Done():
			utils.Log.Info(fmt.Sprintf("Stop saveAndReportVMData goroutine: %#v", uuid))
			return
		case data := <-agent.LibvirtData:
			// report to ws
			err := svc.Dao.WriteVMData(data, map[string]string{
				"uuid": uuid,
			})
			if err != nil {
				utils.Log.Error("Can not write vm data", err)
			}
		}
	}
}

// todo: delete url param
func (svc *Service) StopReport(url, uuid string) error {
	err := agent.StopLibvirtAgent(uuid)
	if err != nil {
		utils.Log.Error("Can not stop agent", err)
		return err
	}

	return nil
}

// get VM metrics
func (svc *Service) GetVMMetrics() []string {
	metrics := make([]string, 0)

	// build VM struct
	data := data.NewLibvirtData(libvirt.DomainCPUStats{}, []libvirt.DomainMemoryStat{}, libvirt.DomainBlockStats{}, libvirt.DomainInterfaceStats{})
	res, err := utils.StructToMapWithJSONTag(data)
	if err != nil {
		utils.Log.Error("Can not convert vm data to map", err)
		return metrics
	}

	// flatten
	flattened := utils.FlattenMap(res, ".")

	// get map keys
	for k := range flattened {
		metrics = append(metrics, k)
	}

	sort.Strings(metrics)

	return metrics
}
