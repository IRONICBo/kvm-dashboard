package main

import (
	"bufio"
	"kvm-dashboard/conf"
	"kvm-dashboard/router"
	"kvm-dashboard/utils"
	"os/exec"

	"github.com/gin-gonic/gin"
)

var config *conf.Config

func SetUpEnv() {
	config = conf.InitConf()
	// xxx
}

func main() {
	cmd := exec.Command("dstat", "-l", "-c", "--proc", "--proc-count", "-y", "-i", "--ipc", "-m", "--vm", "-g", "-s", "-n", "--net-packets", "--socket", "--raw", "--tcp", "--udp", "--unix", "-d", "-r", "--aio", "--disk-tps", "--disk-util", "--fs", "--lock")
	out, err := cmd.StdoutPipe() // start a pipe to get stdout

	if err != nil {
		return
	}

	go func() {
		cmd.Start()
		scanner := bufio.NewScanner(out)

		for scanner.Scan() {
			utils.LogWithInfo(scanner.Text())
		}
	}()

	SetUpEnv()

	gin.SetMode(config.AppConf.Mode)
	r := gin.Default()
	router.WebRouter(r)

	r.Run(config.AppConf.Host + ":" + config.AppConf.Port)
}
