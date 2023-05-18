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
)

func StartVM(c *gin.Context) {
	param := &param.VmArgument{}
	err := c.ShouldBindQuery(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
		return
	}

	go startVM(param.UUID)

	err = ws.ControlWSServer.HandleControlRequest(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
		return
	}
}

func startVM(uuid string) {
	svc := services.NewService()

	// 1
	state := model.NewProgressJson(uuid, 0, "Start vm...", true)
	ws.ControlWSServer.WriteData(uuid, state)
	err := svc.StartVM(uuid)
	if err != nil {
		utils.Log.Error("Can not start vm", err)
		state = model.NewProgressJson(uuid, 100, err.Error(), false)
		ws.ControlWSServer.WriteData(uuid, state)
		ws.ControlWSServer.CloseSession(uuid)
		return
	}

	// 2
	state = model.NewProgressJson(uuid, 60, "Start vm successfully ", true)
	ws.ControlWSServer.WriteData(uuid, state)

	// 3
	svc.StartReport(consts.LIBVIRT_URL, uuid)
	state = model.NewProgressJson(uuid, 100, "Start Report successfully ", true)
	ws.ControlWSServer.WriteData(uuid, state)

	ws.ControlWSServer.CloseSession(uuid)
}

func StopVM(c *gin.Context) {
	param := &param.VmArgument{}
	err := c.ShouldBindQuery(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
		return
	}

	go stopVM(param.UUID)

	err = ws.ControlWSServer.HandleControlRequest(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
		return
	}
}

func stopVM(uuid string) {
	svc := services.NewService()

	// 1
	state := model.NewProgressJson(uuid, 0, "Stop vm...", true)
	ws.ControlWSServer.WriteData(uuid, state)
	err := svc.StopVM(uuid)
	if err != nil {
		utils.Log.Error("Can not stop vm", err)
		state = model.NewProgressJson(uuid, 100, err.Error(), false)
		ws.ControlWSServer.WriteData(uuid, state)
		ws.ControlWSServer.CloseSession(uuid)
		return
	}

	// 2
	state = model.NewProgressJson(uuid, 60, "Stop vm successfully ", true)
	ws.ControlWSServer.WriteData(uuid, state)

	// 3
	err = svc.StopReport(consts.LIBVIRT_URL, uuid)
	if err != nil {
		utils.Log.Error("Can not stop report", err)
		state = model.NewProgressJson(uuid, 100, err.Error(), false)
		ws.ControlWSServer.WriteData(uuid, state)
		ws.ControlWSServer.CloseSession(uuid)
		return
	}

	state = model.NewProgressJson(uuid, 100, "Stop Report successfully ", true)
	ws.ControlWSServer.WriteData(uuid, state)

	ws.ControlWSServer.CloseSession(uuid)
}

func SuspendVM(c *gin.Context) {
	param := &param.VmArgument{}
	err := c.ShouldBindQuery(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
		return
	}

	go suspendVM(param.UUID)

	err = ws.ControlWSServer.HandleControlRequest(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
		return
	}
}

func suspendVM(uuid string) {
	svc := services.NewService()

	// 1
	state := model.NewProgressJson(uuid, 0, "Suspend vm...", true)
	ws.ControlWSServer.WriteData(uuid, state)
	err := svc.SuspendVM(uuid)
	if err != nil {
		utils.Log.Error("Can not suspend vm", err)
		state = model.NewProgressJson(uuid, 100, err.Error(), false)
		ws.ControlWSServer.WriteData(uuid, state)
		ws.ControlWSServer.CloseSession(uuid)
		return
	}

	// 2
	state = model.NewProgressJson(uuid, 60, "Suspend vm successfully ", true)
	ws.ControlWSServer.WriteData(uuid, state)

	// 3
	svc.StopReport(consts.LIBVIRT_URL, uuid)
	state = model.NewProgressJson(uuid, 100, "Stop Report successfully ", true)
	ws.ControlWSServer.WriteData(uuid, state)

	ws.ControlWSServer.CloseSession(uuid)
}

func ResumeVM(c *gin.Context) {
	param := &param.VmArgument{}
	err := c.ShouldBindQuery(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
		return
	}

	go resumeVM(param.UUID)

	err = ws.ControlWSServer.HandleControlRequest(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
		return
	}
}

func resumeVM(uuid string) {
	svc := services.NewService()

	// 1
	state := model.NewProgressJson(uuid, 0, "Resume vm...", true)
	ws.ControlWSServer.WriteData(uuid, state)
	err := svc.ResumeVM(uuid)
	if err != nil {
		utils.Log.Error("Can not resume vm", err)
		state = model.NewProgressJson(uuid, 100, err.Error(), false)
		ws.ControlWSServer.WriteData(uuid, state)
		ws.ControlWSServer.CloseSession(uuid)
		return
	}

	// 2
	state = model.NewProgressJson(uuid, 60, "Resume vm successfully ", true)
	ws.ControlWSServer.WriteData(uuid, state)

	// 3
	svc.StartReport(consts.LIBVIRT_URL, uuid)
	state = model.NewProgressJson(uuid, 100, "Start Report successfully ", true)
	ws.ControlWSServer.WriteData(uuid, state)

	ws.ControlWSServer.CloseSession(uuid)
}
