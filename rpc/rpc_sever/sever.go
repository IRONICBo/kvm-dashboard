package rpc_sever

import (
	"context"
	"fmt"
	"kvm-dashboard/consts"
	"kvm-dashboard/rpc/internal/threshold"
	"kvm-dashboard/utils"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	threshold.UnimplementedRpcSetDomainThvServer
}

func (s *server) SetDomainThv(ctx context.Context, thres *threshold.SetDomainRequest) (*threshold.SetDomainResponse, error) {
	uuid := thres.Uuid
	data, ok := consts.VM_THRESHOLD[uuid]
	if !ok {
		utils.Log.Error(fmt.Sprintf("uuid %s not found", uuid))
		return &threshold.SetDomainResponse{
			Status: false,
		}, fmt.Errorf("uuid %s not found", uuid)
	}

	cpuUsageThreshold, err := strconv.Atoi(thres.CpuThv)
	if err != nil {
		utils.Log.Error(fmt.Sprintf("cpuUsageThreshold %s is not a number", thres.CpuThv))
		return &threshold.SetDomainResponse{
			Status: false,
		}, err
	}
	diskUsageThreshold, err := strconv.Atoi(thres.DiskThv)
	if err != nil {
		utils.Log.Error(fmt.Sprintf("diskUsageThreshold %s is not a number", thres.CpuThv))
		return &threshold.SetDomainResponse{
			Status: false,
		}, err
	}
	memUsageThreshold, err := strconv.Atoi(thres.MemoryThv)
	if err != nil {
		utils.Log.Error(fmt.Sprintf("memUsageThreshold %s is not a number", thres.CpuThv))
		return &threshold.SetDomainResponse{
			Status: false,
		}, err
	}

	data.CpuUsageThreshold = cpuUsageThreshold
	data.DiskUsageThreshold = diskUsageThreshold
	data.MemUsageThreshold = memUsageThreshold
	consts.VM_THRESHOLD[uuid] = data

	resp := &threshold.SetDomainResponse{
		Status: true,
	}
	return resp, nil
}

func StartRpcSever(ip, port string) {
	utils.Log.Info(fmt.Sprintf("Start VMTHresholdServer rpc serve on %s:%s", ip, port))

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", ip, port))
	if err != nil {
		utils.Log.Fatal(fmt.Sprintf("Failed to listen: %v", err))
	}

	s := grpc.NewServer()
	threshold.RegisterRpcSetDomainThvServer(s, &server{})
	reflection.Register(s)

	if s.Serve(lis) != nil {
		utils.Log.Fatal(fmt.Sprintf("Failed to serve: %v", err))
	}
}
