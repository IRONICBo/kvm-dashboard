package services

import (
	"fmt"
	"kvm-dashboard/dao"
	"kvm-dashboard/utils"
	"testing"
)

func TestGetLatestSimpleData(t *testing.T) {
	// Init influxdb
	err := dao.Init(
		"http://127.0.0.1:8086",
		"ldrl4NwtC-Mq0O-N1VmrKN2jh6WpMzA7fXLKpgpfT0tWWZIlf_Doer84jkWcRwgy6cFzVj2ge_HCZFcA_Xfdkg==",
		"hust",
		"kvm-dashboard")
	if err != nil {
		panic(err)
	}

	field := []string{"cpu_usage", "mem_usage", "disk_usage", "net_rx_rate", "net_tx_rate"}
	uuid := "5acefc92-b06d-48c8-a862-889c74d61c23"
	svc := NewService()
	res := svc.GetLatestSimpleData(uuid, field)
	utils.Log.Info(fmt.Sprintf("%#v", res[0]))
}
