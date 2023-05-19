package main

import (
	"encoding/json"
	"kvm-dashboard/conf"
	"kvm-dashboard/dao"
	"kvm-dashboard/router"
	"kvm-dashboard/router/ws"
	"kvm-dashboard/utils"

	"github.com/gin-gonic/gin"
)

var config *conf.Config

func SetUpEnv() {
	config = conf.InitConf()

	configJson, _ := json.MarshalIndent(config, "", "  ")
	utils.Log.Debug("Start with config: \n", string(configJson))

	// Init influxdb
	err := dao.Init(config.InfluxDBConf.URL, config.InfluxDBConf.Token, config.InfluxDBConf.Org, config.InfluxDBConf.Bucket)
	if err != nil {
		panic(err)
	}
}

func main() {
	SetUpEnv()

	gin.SetMode(config.AppConf.Mode)
	r := gin.Default()
	router.WebRouter(r)

	// ws hook
	ws.NewControlWSServer()
	ws.NewProcessWSServer()
	ws.NewSimpleWSServer()

	r.Run(config.AppConf.Host + ":" + config.AppConf.Port)
}
