package agent

import (
	"kvm-dashboard/utils"
	"testing"
	"time"
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

	utils.Log.Info("Start perf agent")

	perfAgent.Start()
	time.Sleep(5 * time.Second)
}
