package api

import (
	"kvm-dashboard/consts"
	"kvm-dashboard/model"
	"kvm-dashboard/router/param"
	"kvm-dashboard/router/ws"
	"kvm-dashboard/services"
	"kvm-dashboard/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

func StartVM(c *gin.Context) {
	param := &param.VmArgument{}
	err := c.ShouldBindQuery(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	m := melody.New()
	m.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	go startVM(param.UUID)

	err = ws.ControlWSServer.HandleControlRequest(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func startVM(uuid string) {
	svc := services.NewService()

	// 1
	model.NewProgressJson(uuid, 0, "Start vm...", true)
	err := svc.StartVM(uuid)
	if err != nil {
		utils.Log.Error("Can not start vm", err)
		state := model.NewProgressJson(uuid, 100, "Can not start vm", false)
		go ws.ControlWSServer.WriteData(uuid, state)
		return
	}

	// 2
	state := model.NewProgressJson(uuid, 60, "Start vm successfully ", true)
	go ws.ControlWSServer.WriteData(uuid, state)

	// 3
	svc.StartReport(consts.LIBVIRT_URL, uuid)
	state = model.NewProgressJson(uuid, 100, "Start Report successfully ", true)
	go ws.ControlWSServer.WriteData(uuid, state)
}

func StopVM(c *gin.Context) {
	param := &param.VmArgument{}
	err := c.ShouldBindQuery(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

}

func SuspendVM(c *gin.Context) {
	param := &param.VmArgument{}
	err := c.ShouldBindQuery(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

}

func ResumeVM(c *gin.Context) {
	param := &param.VmArgument{}
	err := c.ShouldBindQuery(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

}
