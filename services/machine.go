package services

import (
	"context"
	"fmt"
	"kvm-dashboard/consts"
	"kvm-dashboard/model"
	"kvm-dashboard/rpc/api/info"
	"kvm-dashboard/rpc/api/threshold"
	"kvm-dashboard/rpc/rpc_client"
	"kvm-dashboard/utils"
	"strconv"

	"google.golang.org/protobuf/types/known/emptypb"
)

// if we can not get thresold, it will return a default one
func (svc *Service) GetMachineThreshold(uuid string) *model.Threshold {
	if threshold, ok := consts.VM_THRESHOLD[uuid]; ok {
		// if we have threshold in memory, return it
		// the threshold will changed in rpc function
		// so we need not to update it manually
		return threshold
	}

	resp, err := rpc_client.VMTHresholdClient.GetDomainThv(context.Background(), &threshold.GetDomainThvRequest{
		Uuid: uuid,
	})
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Failed to get machine threshold: %v", err))
		defaultThreshold := model.NewThreshold(consts.CPU_USAGE_THRESHOLD, consts.MEM_USAGE_THRESHOLD, consts.DISK_USAGE_THRESHOLD)
		consts.VM_THRESHOLD[uuid] = defaultThreshold

		return defaultThreshold
	}

	var cpuUsageThreshold, memUsageThreshold, diskUsageThreshold int
	// cpu threshold
	res, err := strconv.ParseFloat(resp.CpuThv, 64) // 0.8
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Failed to convert cpu threshold: %v", err))
		cpuUsageThreshold = consts.CPU_USAGE_THRESHOLD
	} else {
		res = res * 100
		cpuUsageThreshold = int(res)
	}

	// mem threshold
	res, err = strconv.ParseFloat(resp.MemoryThv, 64) // 0.8
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Failed to convert mem threshold: %v", err))
		memUsageThreshold = consts.MEM_USAGE_THRESHOLD
	} else {
		res = res * 100
		memUsageThreshold = int(res)
	}

	// disk threshold
	res, err = strconv.ParseFloat(resp.DiskThv, 64) // 0.8
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Failed to convert disk threshold: %v", err))
		diskUsageThreshold = consts.DISK_USAGE_THRESHOLD
	} else {
		res = res * 100
		diskUsageThreshold = int(res)
	}

	threshold := model.NewThreshold(cpuUsageThreshold, memUsageThreshold, diskUsageThreshold)
	consts.VM_THRESHOLD[uuid] = threshold
	return threshold
}

func (svc *Service) GetMachineList() (*model.MachineInfo, error) {
	resp, err := rpc_client.MachineInfoClient.GetMachineInfo(context.Background(), &emptypb.Empty{})
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Failed to get machine info: %v", err))
		return model.NewMachineInfo(
			make([]*model.Machine, 0),
			make([]*model.Machine, 0)), err
	}

	hostInfos := resp.HostInfo
	vmInfos := resp.VmInfo

	// temp hostinfo map for vm
	hostInfoMap := make(map[string]*info.HostInfo) // machine info in grpc

	hosts := make([]*model.Machine, 0)
	for _, hostInfo := range hostInfos {
		host := model.NewMachine(
			hostInfo.Name,
			hostInfo.Uuid,
			hostInfo.Ip,
			hostInfo.SshPort,
			fmt.Sprintf("qemu+tcp://%s:%s/system", hostInfo.Ip, hostInfo.QemuPort),
			hostInfo.Username,
			hostInfo.Passwd,
		)
		hosts = append(hosts, host)

		// temp map
		hostInfoMap[hostInfo.Uuid] = hostInfo
	}

	vms := make([]*model.Machine, 0)
	for _, vmInfo := range vmInfos {
		hostInfo := hostInfoMap[vmInfo.HostId]

		vm := model.NewMachine(
			vmInfo.Name,
			vmInfo.Uuid,
			"",
			fmt.Sprintf("%d", consts.PORT),
			fmt.Sprintf("qemu+tcp://%s:%s/system", hostInfo.Ip, hostInfo.QemuPort),
			consts.USERNAME,
			consts.PASSWORD,
		)
		vms = append(vms, vm)
	}

	machineInfo := model.NewMachineInfo(hosts, vms)
	return machineInfo, nil
}
