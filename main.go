package main

import (
	"kvm-dashboard/conf"
	"kvm-dashboard/router"

	"github.com/gin-gonic/gin"
)

var config *conf.Config

func SetUpEnv() {
	config = conf.InitConf()
	// xxx
}

func main() {
	SetUpEnv()

	gin.SetMode(config.AppConf.Mode)
	r := gin.Default()
	router.WebRouter(r)

	r.Run(config.AppConf.Host + ":" + config.AppConf.Port)
}
