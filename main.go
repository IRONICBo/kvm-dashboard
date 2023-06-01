package main

import (
	"encoding/json"
	"kvm-dashboard/conf"
	"kvm-dashboard/dao"
	"kvm-dashboard/middleware"
	"kvm-dashboard/router"
	"kvm-dashboard/router/ws"
	"kvm-dashboard/rpc/rpc_client"
	"kvm-dashboard/rpc/rpc_sever"
	"kvm-dashboard/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetUpEnv() {
	conf.C = conf.InitConf()

	configJson, _ := json.MarshalIndent(conf.C, "", "  ")
	utils.Log.Debug("Start with config: \n", string(configJson))

	// Init influxdb
	err := dao.Init(conf.C.InfluxDBConf.URL, conf.C.InfluxDBConf.Token, conf.C.InfluxDBConf.Org, conf.C.InfluxDBConf.Bucket)
	if err != nil {
		panic(err)
	}
}

func main() {
	SetUpEnv()

	gin.SetMode(conf.C.AppConf.Mode)
	r := gin.Default()
	store := cookie.NewStore([]byte("asklv"))
	r.Use(sessions.Sessions("asklv-sessions", store)) // use session middleware
	r.Use(middleware.EnableCROS())
	router.WebRouter(r)

	// ws hook
	ws.NewControlWSServer()
	ws.NewProcessWSServer()
	ws.NewSimpleWSServer()

	// rpc
	go rpc_sever.StartRpcSever(conf.C.RpcConf.SeverHost, conf.C.RpcConf.SeverPort)
	go rpc_client.StartRpcClient(conf.C.RpcConf.ClientHost, conf.C.RpcConf.ClientPort)

	r.Run(conf.C.AppConf.Host + ":" + conf.C.AppConf.Port)
}
