package services

import (
	"fmt"
	"kvm-dashboard/dao"
	"kvm-dashboard/utils"
	"testing"
	"time"
)

func TestStartVMReport(t *testing.T) {
	// Init influxdb
	err := dao.Init(
		"http://127.0.0.1:8086",
		"ldrl4NwtC-Mq0O-N1VmrKN2jh6WpMzA7fXLKpgpfT0tWWZIlf_Doer84jkWcRwgy6cFzVj2ge_HCZFcA_Xfdkg==",
		"hust",
		"kvm-dashboard")
	if err != nil {
		panic(err)
	}

	svc := NewService()
	svc.StartVMReport("qemu:///system", "a5c655a8-589a-4524-b338-fa9b947a334d")

	time.Sleep(60 * time.Second)
}

func TestGetLatestVMData(t *testing.T) {
	// Init influxdb
	err := dao.Init(
		"http://127.0.0.1:8086",
		"ldrl4NwtC-Mq0O-N1VmrKN2jh6WpMzA7fXLKpgpfT0tWWZIlf_Doer84jkWcRwgy6cFzVj2ge_HCZFcA_Xfdkg==",
		"hust",
		"kvm-dashboard")
	if err != nil {
		panic(err)
	}

	field := []string{"interface_stats.rx_packets"}
	uuid := "a5c655a8-589a-4524-b338-fa9b947a334d"
	svc := NewService()
	res := svc.GetLatestVMData(uuid, field)
	utils.Log.Info(fmt.Sprintf("%#v", res[0]))
}

func TestGetVMMetrics(t *testing.T) {
	svc := NewService()
	res := svc.GetVMMetrics()
	utils.Log.Info(fmt.Sprintf("%#v", res))
}
