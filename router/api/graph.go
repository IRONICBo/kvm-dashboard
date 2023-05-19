package api

import (
	"kvm-dashboard/model"
	"kvm-dashboard/router/param"
	"kvm-dashboard/router/ws"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHostWorkloadHistoryStats(c *gin.Context) {

}

func GetHostWorkloadRealtimeStats(c *gin.Context) {

}

func GetVmWorkloadHistoryStats(c *gin.Context) {
	// get history data
}

func GetVmWorkloadRealtimeStats(c *gin.Context) {
	// get realtime data
	param := &param.VmArgument{}
	err := c.ShouldBindQuery(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
		return
	}

	// init websocket result
	err = ws.SimpleWSServer.HandleSimpleRequest(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
		return
	}
}
