package services

import (
	"context"
	"fmt"
	"kvm-dashboard/model"
	"kvm-dashboard/rpc/rpc_client"
	"kvm-dashboard/utils"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (svc *Service) GetMachineList() {
	// &emptypb.Empty{} => empty struct
	resp, err := rpc_client.MachineInfoClient.GetMachineInfo(context.Background(), &emptypb.Empty{})
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Failed to get machine info: %v", err))
	}

	respHostList := resp.HostInfo
	hostInfoList := make([]*model.HostInfo, 0)
	for _, hostInfo := range respHostList {
		model.NewHost(
			hostInfo.Name,
			hostInfo.Uuid,
			hostInfo.Ip,
			hostInfo.SshPort,
			hostInfo.Username,
			host
		)
		utils.Log.Info(fmt.Sprintf("hostInfo %v", hostInfo))
	}
}
