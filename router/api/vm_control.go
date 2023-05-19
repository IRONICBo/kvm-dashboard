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

func StartReport(c *gin.Context) {
	param := &param.VmArgument{}
	err := c.ShouldBindQuery(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
		return
	}

	go startAllReport(services.NewService(), param.UUID)

	err = ws.ControlWSServer.HandleControlRequest(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
		return
	}
}

func StopReport(c *gin.Context) {
	param := &param.VmArgument{}
	err := c.ShouldBindQuery(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
		return
	}

	go stopAllReport(services.NewService(), param.UUID)

	err = ws.ControlWSServer.HandleControlRequest(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
		return
	}
}

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
	// err := svc.StartVM(uuid)
	// if err != nil {
	// 	utils.Log.Error("Can not start vm", err)
	// 	state = model.NewProgressJson(uuid, 100, err.Error(), false)
	// 	ws.ControlWSServer.WriteData(uuid, state)
	// 	ws.ControlWSServer.CloseSession(uuid)
	// 	return
	// }

	// 2
	state = model.NewProgressJson(uuid, 50, "Start vm successfully ", true)
	ws.ControlWSServer.WriteData(uuid, state)

	startAllReport(svc, uuid)

	ws.ControlWSServer.CloseSession(uuid)
}

func startAllReport(svc *services.Service, uuid string) {
	// 2
	state := model.NewProgressJson(uuid, 60, "Start perf report successfully ", true)
	ws.ControlWSServer.WriteData(uuid, state)

	// 3
	err := svc.StartVMReport(consts.LIBVIRT_URL, uuid)
	if err != nil {
		utils.Log.Error("Can not start vm report", err)
		state := model.NewProgressJson(uuid, 100, err.Error(), false)
		ws.ControlWSServer.WriteData(uuid, state)
		ws.ControlWSServer.CloseSession(uuid)
		return
	}
	state = model.NewProgressJson(uuid, 70, "Start vm report successfully ", true)
	ws.ControlWSServer.WriteData(uuid, state)

	// 4
	err = svc.StartProcessReport(uuid)
	if err != nil {
		utils.Log.Error("Can not start process report", err)
		state := model.NewProgressJson(uuid, 100, err.Error(), false)
		ws.ControlWSServer.WriteData(uuid, state)
		ws.ControlWSServer.CloseSession(uuid)
		return
	}
	state = model.NewProgressJson(uuid, 80, "Start process info successfully ", true)
	ws.ControlWSServer.WriteData(uuid, state)

	// 5
	err = svc.StartSimpleReport(uuid)
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
	// err := svc.StopVM(uuid)
	// if err != nil {
	// 	utils.Log.Error("Can not stop vm", err)
	// 	state = model.NewProgressJson(uuid, 100, err.Error(), false)
	// 	ws.ControlWSServer.WriteData(uuid, state)
	// 	ws.ControlWSServer.CloseSession(uuid)
	// 	return
	// }

	state = model.NewProgressJson(uuid, 50, "Stop vm successfully ", true)
	ws.ControlWSServer.WriteData(uuid, state)

	stopAllReport(svc, uuid)

	ws.ControlWSServer.CloseSession(uuid)
}

func stopAllReport(svc *services.Service, uuid string) {
	// 2
	state := model.NewProgressJson(uuid, 60, "Stop perf report successfully ", true)
	ws.ControlWSServer.WriteData(uuid, state)

	// 3
	err := svc.StopVMReport(consts.LIBVIRT_URL, uuid)
	if err != nil {
		utils.Log.Error("Can not stop vm report", err)
		state := model.NewProgressJson(uuid, 100, err.Error(), false)
		ws.ControlWSServer.WriteData(uuid, state)
		ws.ControlWSServer.CloseSession(uuid)
		return
	}
	state = model.NewProgressJson(uuid, 70, "Stop vm report successfully ", true)
	ws.ControlWSServer.WriteData(uuid, state)

	// 4
	err = svc.StopProcessReport(uuid)
	if err != nil {
		utils.Log.Error("Can not stop process report", err)
		state = model.NewProgressJson(uuid, 100, err.Error(), false)
		ws.ControlWSServer.WriteData(uuid, state)
		ws.ControlWSServer.CloseSession(uuid)
		return
	}
	state = model.NewProgressJson(uuid, 80, "Stop process info successfully ", true)
	ws.ControlWSServer.WriteData(uuid, state)

	// 5
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

	state = model.NewProgressJson(uuid, 50, "Suspend vm successfully ", true)
	ws.ControlWSServer.WriteData(uuid, state)

	stopAllReport(svc, uuid)

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

	state = model.NewProgressJson(uuid, 50, "Resume vm successfully ", true)
	ws.ControlWSServer.WriteData(uuid, state)

	startAllReport(svc, uuid)

	ws.ControlWSServer.CloseSession(uuid)
}
