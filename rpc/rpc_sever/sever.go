package rpc_sever

import (
	"context"
	"fmt"
	"kvm-dashboard/consts"
	"kvm-dashboard/rpc/api/threshold"
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

		// need not to modify threshold in memory
		return &threshold.SetDomainResponse{
			Status: false,
		}, nil
	}

	var cpuUsageThreshold, memUsageThreshold, diskUsageThreshold int
	// cpu threshold
	res, err := strconv.ParseFloat(thres.CpuThv, 64) // 0.8
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Failed to convert cpu threshold: %v", err))
		cpuUsageThreshold = consts.CPU_USAGE_THRESHOLD
	} else {
		res = res * 100
		cpuUsageThreshold = int(res)
	}

	// mem threshold
	res, err = strconv.ParseFloat(thres.MemoryThv, 64) // 0.8
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Failed to convert mem threshold: %v", err))
		memUsageThreshold = consts.MEM_USAGE_THRESHOLD
	} else {
		res = res * 100
		memUsageThreshold = int(res)
	}

	// disk threshold
	res, err = strconv.ParseFloat(thres.DiskThv, 64) // 0.8
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Failed to convert disk threshold: %v", err))
		diskUsageThreshold = consts.DISK_USAGE_THRESHOLD
	} else {
		res = res * 100
		diskUsageThreshold = int(res)
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
