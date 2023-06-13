package router

import (
	"kvm-dashboard/middleware"
	"kvm-dashboard/router/api"

	"github.com/gin-gonic/gin"
)

func WebRouter(r *gin.Engine) {
	kvm_dashboard_api := r.Group("/api-go")

	UserInfo := kvm_dashboard_api.Group("/user")
	{
		UserInfo.POST("/login", api.Login)
		UserInfo.GET("/logout", api.Logout)
	}

	MachineInfo := kvm_dashboard_api.Group("/machine")
	MachineInfo.Use(middleware.CheckValidTimestamp())
	{
		MachineInfo.GET("/list", api.GetMachineList)
		MachineInfo.GET("/threshold", api.GetMachineThreshold)
	}

	HostInfo := kvm_dashboard_api.Group("/host")
	HostInfo.Use(middleware.CheckValidTimestamp())
	{
		HostInfo.GET("/basic", api.GetHostBasicInfo)

		Device := HostInfo.Group("/device")
		{
			Device.GET("/disk/list", api.GetHostDiskList)
			Device.GET("/interface/list", api.GetHostInterfaceList)
		}

		Control := HostInfo.Group("/control")
		{
			Control.GET("/start_report/ws", api.StartHostReport)
			Control.GET("/stop_report/ws", api.StopHostReport)
		}

		Graph := HostInfo.Group("/graph")
		{
			Graph.POST("/workload/history", api.GetHostWorkloadHistoryStats)
			Graph.GET("/workload/realtime/ws", api.GetHostWorkloadRealtimeStats)
		}
		Table := HostInfo.Group("/table")
		{
			Table.GET("/process/realtime/ws", api.GetHostProcessRealtimeStats)

			Table.POST("/alert/history", api.GetHostProcessHistoryAlertInfo)
			Table.GET("/alert/realtime/ws", api.GetHostProcessRealtimeAlertInfo)
		}
	}

	VmInfo := kvm_dashboard_api.Group("/vm")
	VmInfo.Use(middleware.CheckValidTimestamp())
	{
		VmInfo.GET("/basic", api.GetVmBasicInfo)

		Device := HostInfo.Group("/device")
		{
			Device.GET("/disk/list", api.GetVMDiskList)
			Device.GET("/interface/list", api.GetVMInterfaceList)
		}

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
