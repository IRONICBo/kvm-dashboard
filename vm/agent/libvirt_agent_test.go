package agent

import (
	"encoding/json"
	"fmt"
	"kvm-dashboard/utils"
	"testing"
	"time"
)

func TestInitVm(t *testing.T) {
	libvirtAgent, _ := NewLibvirtAgent(
		"qemu:///system",
	)

	uuid := "917a29bb-ba86-484b-b8d3-c80c27dddc36"
	err := libvirtAgent.VmSetUp(uuid, "root", "123123")
	utils.Log.Info(err.Error())
}

func TestLibvirtAgent(t *testing.T) {
	libvirtAgent, _ := NewLibvirtAgent(
		// &AgentInfo{
		// 	UUID: "a5c655a8-589a-4524-b338-fa9b947a334d",
		// 	// 917a29bb-ba86-484b-b8d3-c80c27dddc36
		// 	User:     "root",
		// 	Password: "123123",
		// 	Ip:       "192.168.122.166",
		// 	Port:     22,
		// },
		"qemu:///system",
	)

	uuid := "a5c655a8-589a-4524-b338-fa9b947a334d"
	utils.Log.Info("Start libvirt agent")
	// libvirtAgent.Start(uuid) // start vm

	info, _ := libvirtAgent.GetVMInfo(uuid)
	utils.Log.Info(fmt.Sprintf("%#v", info))
	utils.Log.Info(info.MaxMem/1024/1024, "GB")

	for {
		select {
		case <-time.After(10 * time.Second):
			utils.Log.Info("Time out")
			return
		case data := <-libvirtAgent.LibvirtData:
			// if !ok {
			// 	utils.Log.Info("Channel closed")
			// 	return
			// }
			// 转换为json输出
			jsonData, err := json.Marshal(data)
			if err != nil {
				utils.Log.Error(fmt.Sprintf("Can not marshal libvirtData: %#v", uuid), err)
				return
			}
			utils.Log.Info(string(jsonData))
		}
	}

	// libvirtAgent.StartVM(uuid)
	// time.Sleep(5 * time.Second)
	// libvirtAgent.StopVM(uuid)
	// time.Sleep(5 * time.Second)
	// libvirtAgent.SuspendVM(uuid)
	// time.Sleep(5 * time.Second)
	// libvirtAgent.ResumeVM(uuid)
	// time.Sleep(5 * time.Second)

}

// {
// 	"cpu_stats": {
// 		"cpu_time": 225018619810,
// 		"user_time": 319590000000,
// 		"system_time": 29940000000,
// 		"vcpu_time": 0
// 	},
// 	"memory_stats": {
// 		"swap_in": 0,
// 		"swap_out": 0,
// 		"major_fault": 0,
// 		"minor_fault": 0,
// 		"unused": 0,
// 		"available": 0,
// 		"actual_ballon": 0,
// 		"rss": 1775104,
// 		"usable": 0,
// 		"last_update": 0,
// 		"disk_caches": 0,
// 		"huge_tlb_pg_alloc": 0,
// 		"huge_tlb_pg_fail": 0
// 	},
// 	"block_stats": {
// 		"rd_bytes": 640500224,
// 		"rd_req": 7528,
// 		"rd_total_times": 0,
// 		"wr_bytes": 96374272,
// 		"wr_req": 6706,
// 		"wr_total_times": 0,
// 		"flush_req": 0,
// 		"flush_total_times": 0,
// 		"errs": 0
// 	},
// 	"interface_stats": {
// 		"rx_bytes": 748782,
// 		"rx_packets": 14303,
// 		"rx_errs": 0,
// 		"rx_drop": 0,
// 		"tx_bytes": 16914,
// 		"tx_packets": 223,
// 		"tx_errs": 0,
// 		"tx_drop": 0
// 	}
// }
