package api

import (
	"kvm-dashboard/model"
	"kvm-dashboard/router/param"
	"kvm-dashboard/router/ws"
	"kvm-dashboard/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHostWorkloadHistoryStats(c *gin.Context) {
	// get history data
	param := &param.HostHistoryArgument{}
	err := c.ShouldBind(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
		return
	}

	if len(param.Fields) == 0 {
		// fix fields
		param.Fields = []string{
			"cpu_usage",
			"mem_usage",
			"disk_usage",
			"net_rx_rate",
			"net_tx_rate",
		}
	}

	svc := services.NewService()
	hostHistorySimpleData := svc.GetLatestSimpleData(param.UUID, param.Fields)

	c.JSON(http.StatusOK, model.SuccessWithData(hostHistorySimpleData))
}

func GetHostWorkloadRealtimeStats(c *gin.Context) {
	// get realtime data
	param := &param.HostArgument{}
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

func GetMetricHistoryStats(c *gin.Context) {
	param := &param.VmMetricArgument{}
	err := c.ShouldBind(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
		return
	}

	svc := services.NewService()
	metricHistoryData := svc.GetVMData(param.UUID, param.Period, param.Func, param.Fields)

	c.JSON(http.StatusOK, model.SuccessWithData(metricHistoryData))
}

func GetMetricList(c *gin.Context) {
	svc := services.NewService()
	metricList := svc.GetVMMetrics()

	c.JSON(http.StatusOK, model.SuccessWithData(metricList))
}

func GetVmWorkloadHistoryStats(c *gin.Context) {
	// get history data
	param := &param.VmHistoryArgument{}
	err := c.ShouldBind(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
		return
	}

	if len(param.Fields) == 0 {
		// fix fields
		param.Fields = []string{
			"cpu_usage",
			"mem_usage",
			"disk_usage",
			"net_rx_rate",
			"net_tx_rate",
		}
	}

	svc := services.NewService()
	vmHistorySimpleData := svc.GetLatestSimpleData(param.UUID, param.Fields)

	c.JSON(http.StatusOK, model.SuccessWithData(vmHistorySimpleData))
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
