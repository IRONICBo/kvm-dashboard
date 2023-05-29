package rpc_client

import (
	"fmt"
	"kvm-dashboard/rpc/api/info"
	"kvm-dashboard/rpc/api/threshold"
	"kvm-dashboard/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var RpcConn *Conn
var MachineInfoClient info.RpcMachineInfoClient
var VMTHresholdClient threshold.RpcDomainThvClient

type Conn struct {
	conn *grpc.ClientConn
}

func StartRpcClient(ip, port string) *Conn {
	utils.Log.Info(fmt.Sprintf("Start rpc client on %s:%s", ip, port))
	conn, err := grpc.Dial(ip+":"+port, grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))

	if err != nil {
		utils.Log.Fatal(fmt.Sprintf("Failed to dial: %v", err))
	}

	// init client
	RpcConn = &Conn{conn: conn}
	MachineInfoClient = info.NewRpcMachineInfoClient(conn)
	VMTHresholdClient = threshold.NewRpcDomainThvClient(conn)

	return RpcConn
}

func (c *Conn) Close() {
	c.conn.Close()
}
