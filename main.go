package main

import (
	"kvm-dashboard/conf"
	"kvm-dashboard/dao"
	"kvm-dashboard/router"
	"kvm-dashboard/router/ws"

	"github.com/gin-gonic/gin"
)

var config *conf.Config

func SetUpEnv() {
	config = conf.InitConf()

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

	r.Run(config.AppConf.Host + ":" + config.AppConf.Port)
}
