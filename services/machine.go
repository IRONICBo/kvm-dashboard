package services

import (
	"context"
	"fmt"
	"kvm-dashboard/consts"
	"kvm-dashboard/model"
	"kvm-dashboard/rpc/api/info"
	"kvm-dashboard/rpc/rpc_client"
	"kvm-dashboard/utils"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (svc *Service) GetMachineList() (*model.MachineInfo, error) {
	resp, err := rpc_client.MachineInfoClient.GetMachineInfo(context.Background(), &emptypb.Empty{})
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Failed to get machine info: %v", err))
		return model.NewMachineInfo(
			make([]*model.Host, 0),
			make([]*model.Vm, 0)), err
	}

	hostInfos := resp.HostInfo
	vmInfos := resp.VmInfo

	// temp hostinfo map for vm
	hostInfoMap := make(map[string]*info.HostInfo) // machine info in grpc

	hosts := make([]*model.Host, 0)
	for _, hostInfo := range hostInfos {
		host := model.NewHost(
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

	vms := make([]*model.Vm, 0)
	for _, vmInfo := range vmInfos {
		hostInfo := hostInfoMap[vmInfo.HostId]

		vm := model.NewVm(
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
