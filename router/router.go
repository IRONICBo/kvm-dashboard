package router

import (
	"kvm-dashboard/router/api"

	"github.com/gin-gonic/gin"
)

func WebRouter(r *gin.Engine) {
	kvm_dashboard_api := r.Group("/api")

	HosInfo := kvm_dashboard_api.Group("/host")
	{
		HosInfo.GET("/basic", api.GetHostBasicInfo)
		Graph := HosInfo.Group("/graph")
		{
			Graph.POST("/workload/history", api.GetHostWorkloadHistoryStats)
			Graph.GET("/workload/realtime/ws", api.GetHostWorkloadRealtimeStats)
		}
		Table := HosInfo.Group("/table")
		{
			Table.GET("/process/realtime/ws", api.GetHostProcessRealtimeStats)

			Table.POST("/alert/history", api.GetHostProcessHistoryAlertInfo)
			Table.GET("/alert/realtime/ws", api.GetHostProcessRealtimeAlertInfo)
		}
	}

	VmInfo := kvm_dashboard_api.Group("/vm")
	{
		VmInfo.GET("/basic", api.GetVmBasicInfo)

		Control := VmInfo.Group("/control")
		{
			Control.GET("/start/ws", api.StartVM)
			Control.GET("/stop/ws", api.StopVM)
			Control.GET("/suspend/ws", api.SuspendVM)
			Control.GET("/resume/ws", api.ResumeVM)
			Control.GET("/start_report/ws", api.StartReport)
			Control.GET("/stop_report/ws", api.StopReport)
		}

		Graph := VmInfo.Group("/graph")
		{
			Graph.POST("/workload/history", api.GetVmWorkloadHistoryStats)
			Graph.GET("/workload/realtime/ws", api.GetVmWorkloadRealtimeStats)

			Graph.POST("/metric/history", api.GetMetricHistoryStats)
			Graph.GET("/metric/list", api.GetMetricList)
		}

		Table := VmInfo.Group("/table")
		{
			Table.GET("/process/realtime/ws", api.GetVmProcessRealtimeStats)

			Table.POST("/alert/history", api.GetVmAlertHistoryAlertInfo)
			// deprecated
			// Table.GET("/alert/realtime/ws", api.GetVmAlertRealtimeAlertInfo)
		}
	}

}
