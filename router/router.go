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
			Graph.GET("/workload/history", api.GetHostWorkloadHistoryStats)
			Graph.GET("/workload/realtime/ws", api.GetHostWorkloadRealtimeStats)
		}
		Table := HosInfo.Group("/table")
		{
			Table.GET("/process/realtime/ws", api.GetHostProcessRealtimeStats)

			Table.GET("/alert/history", api.GetHostProcessHistoryAlertInfo)
			Table.GET("/alert/realtime/ws", api.GetHostProcessRealtimeAlertInfo)
		}
	}

	VmInfo := kvm_dashboard_api.Group("/vm")
	{
		VmInfo.GET("/basic", api.GetVmBasicInfo)
		Graph := VmInfo.Group("/graph")
		{
			Graph.GET("/workload/history", api.GetVmWorkloadHistoryStats)
			Graph.GET("/workload/realtime/ws", api.GetVmWorkloadRealtimeStats)
		}
		Table := VmInfo.Group("/table")
		{
			Table.GET("/process/realtime/ws", api.GetVmProcessRealtimeStats)

			Table.GET("/alert/history", api.GetVmAlertHistoryAlertInfo)
			Table.GET("/alert/realtime/ws", api.GetVmAlertRealtimeAlertInfo)
		}
	}
}
