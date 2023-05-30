package api

import (
	"kvm-dashboard/model"
	"kvm-dashboard/router/param"
	"kvm-dashboard/router/ws"
	"kvm-dashboard/services"
	"kvm-dashboard/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartHostReport(c *gin.Context) {
	param := &param.VmArgument{}
	err := c.ShouldBindQuery(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
		return
	}

	go startHostReport(services.NewService(), param.UUID)

	err = ws.ControlWSServer.HandleControlRequest(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
		return
	}
}

func StopHostReport(c *gin.Context) {
	param := &param.VmArgument{}
	err := c.ShouldBindQuery(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
		return
	}

	go stopHostReport(services.NewService(), param.UUID)

	err = ws.ControlWSServer.HandleControlRequest(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
		return
	}
}

func startHostReport(svc *services.Service, uuid string) {
	// 1
	state := model.NewProgressJson(uuid, 60, "Start perf report successfully ", true)
	ws.ControlWSServer.WriteData(uuid, state)

	// 2
	err := svc.StartHostProcessReport(uuid)
	if err != nil {
		utils.Log.Error("Can not start process report", err)
		state := model.NewProgressJson(uuid, 100, err.Error(), false)
		ws.ControlWSServer.WriteData(uuid, state)
		ws.ControlWSServer.CloseSession(uuid)
		return
	}
	state = model.NewProgressJson(uuid, 80, "Start process info successfully ", true)
	ws.ControlWSServer.WriteData(uuid, state)

	// 3
	err = svc.StartHostSimpleReport(uuid)
	if err != nil {
		utils.Log.Error("Can not start simple report", err)
		state := model.NewProgressJson(uuid, 100, err.Error(), false)
		ws.ControlWSServer.WriteData(uuid, state)
		ws.ControlWSServer.CloseSession(uuid)
		return
	}
	state = model.NewProgressJson(uuid, 90, "Start simple info successfully ", true)
	ws.ControlWSServer.WriteData(uuid, state)

	// 6
	// err  = svc.StartDstatReport(uuid)
	// if err != nil {
	// 	utils.Log.Error("Can not start dstat report", err)
	// 	state := model.NewProgressJson(uuid, 100, err.Error(), false)
	// 	ws.ControlWSServer.WriteData(uuid, state)
	// 	ws.ControlWSServer.CloseSession(uuid)
	// 	return
	// }
	state = model.NewProgressJson(uuid, 100, "Start dstat info successfully ", true)
	ws.ControlWSServer.WriteData(uuid, state)
}

func stopHostReport(svc *services.Service, uuid string) {
	// 1
	state := model.NewProgressJson(uuid, 60, "Stop perf report successfully ", true)
	ws.ControlWSServer.WriteData(uuid, state)

	// 2
	err := svc.StopProcessReport(uuid)
	if err != nil {
		utils.Log.Error("Can not stop process report", err)
		state = model.NewProgressJson(uuid, 100, err.Error(), false)
		ws.ControlWSServer.WriteData(uuid, state)
		ws.ControlWSServer.CloseSession(uuid)
		return
	}
	state = model.NewProgressJson(uuid, 80, "Stop process info successfully ", true)
	ws.ControlWSServer.WriteData(uuid, state)

	// 3
	err = svc.StopSimpleReport(uuid)
	if err != nil {
		utils.Log.Error("Can not stop simple report", err)
		state = model.NewProgressJson(uuid, 100, err.Error(), false)
		ws.ControlWSServer.WriteData(uuid, state)
		ws.ControlWSServer.CloseSession(uuid)
		return
	}
	state = model.NewProgressJson(uuid, 90, "Stop simple info successfully ", true)
	ws.ControlWSServer.WriteData(uuid, state)

	// 6
	// err = svc.StopDstatReport(uuid)
	// if err != nil {
	// 	utils.Log.Error("Can not stop dstat info", err)
	// 	state = model.NewProgressJson(uuid, 100, err.Error(), false)
	// 	ws.ControlWSServer.WriteData(uuid, state)
	// 	ws.ControlWSServer.CloseSession(uuid)
	// 	return
	// }
	state = model.NewProgressJson(uuid, 100, "Stop dstat info successfully ", true)
	ws.ControlWSServer.WriteData(uuid, state)
}
