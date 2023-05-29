package services

import (
	"context"
	"fmt"
	"kvm-dashboard/rpc/api/threshold"
	"kvm-dashboard/rpc/rpc_client"
	"kvm-dashboard/utils"
)

func (svc *Service) GetMachineList() {
	// &emptypb.Empty{} => empty struct
	// resp, err := rpc_client.MachineInfoClient.GetMachineInfo(context.Background(), &emptypb.Empty{})
	req := &threshold.GetDomainThvRequest{
		Uuid: "a5c655a8-589a-4524-b338-fa9b947a334d",
	}
	resp, err := rpc_client.VMTHresholdClient.GetDomainThv(context.Background(), req)
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Failed to get machine info: %v", err))
	}

	utils.Log.Info(fmt.Sprintf("resp %v", resp))
	// resp = resp
	// respHostList := resp.HostInfo
	// hostInfoList := make([]*model.HostInfo, 0)
	// for _, hostInfo := range respHostList {
	// 	model.NewHost(
	// 		hostInfo.Name,
	// 		hostInfo.Uuid,
	// 		hostInfo.Ip,
	// 		hostInfo.SshPort,
	// 		hostInfo.Username,
	// 	)
	// 	utils.Log.Info(fmt.Sprintf("hostInfo %v", hostInfo))
	// }
}
