package agent

import (
	"kvm-dashboard/utils"
	"testing"
)

func TestLibvirtAgent(t *testing.T) {
	libvirtAgent := &LibvirtAgent{
		AgentInfo: &AgentInfo{
			UUID:     "a5c655a8-589a-4524-b338-fa9b947a334d",
			User:     "root",
			Password: "123123",
			Ip:       "192.168.122.166",
			Port:     22,
		},
		URL: "qemu:///system",
	}

	utils.LogWithInfo("Start libvirt agent")
	libvirtAgent.Start()
}
