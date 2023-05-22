package api

import (
	"kvm-dashboard/model"
	"kvm-dashboard/router/param"
	"kvm-dashboard/router/ws"
	"kvm-dashboard/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHostProcessRealtimeStats(c *gin.Context) {

}

func GetHostProcessHistoryAlertInfo(c *gin.Context) {

}

func GetHostProcessRealtimeAlertInfo(c *gin.Context) {

}

func GetVmProcessRealtimeStats(c *gin.Context) {
	param := &param.VmArgument{}
	err := c.ShouldBindQuery(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
		return
	}

	// fix start multi process
	// svc := services.NewService()
	// go svc.StartProcessReport(param.UUID)

	err = ws.ProcessWSServer.HandleControlRequest(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
		return
	}
}

func GetVmAlertHistoryAlertInfo(c *gin.Context) {
	param := &param.VmPageAgrument{}
	err := c.ShouldBind(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
		return
	}

	svc := services.NewService()
	pageData := svc.GetAlertInfo(param.UUID, param.PageSize, param.Page)

	c.JSON(http.StatusOK, model.SuccessWithData(pageData))
}

func GetVmAlertRealtimeAlertInfo(c *gin.Context) {

}
