package agent

import (
	"kvm-dashboard/utils"
	"testing"
)

func TestPerfAgent(t *testing.T) {
	perfAgent := &PerfAgent{
		AgentInfo: &AgentInfo{
			UUID:     "test",
			User:     "root",
			Password: "123123",
			Ip:       "192.168.122.166",
			Port:     22,
		},
	}

	utils.LogWithInfo("Start perf agent")

	perfAgent.Start()
	// time.Sleep(10000000 * time.Second)
}
