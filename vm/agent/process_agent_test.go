package agent

import (
	"kvm-dashboard/utils"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

func TestProcessAgent(t *testing.T) {
	processAgent, err := NewProcessAgent(
		&AgentInfo{
			UUID:     "a5c655a8-589a-4524-b338-fa9b947a334d",
			User:     "root",
			Password: "123123",
			Ip:       "192.168.122.166",
			Port:     22,
		},
	)
	assert.Equal(t, err, nil)

	utils.Log.Info("Start dstat agent")

	uuid := "a5c655a8-589a-4524-b338-fa9b947a334d"
	processAgent.Start(uuid)
	time.Sleep(10 * time.Second)
}
