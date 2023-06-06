package vm_utils

import (
	"kvm-dashboard/utils"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestJumpToServer(t *testing.T) {
	base, err := NewJumpServer(
		"root",
		"",
		"127.0.0.1",
		22,
	)
	assert.Equal(t, err, nil)

	client, err := JumpToServer(
		base,
		"root",
		"",
		"192.168.122.166",
		22,
	)
	assert.Equal(t, err, nil)

	output, _ := GetCommandOutput(client, "free", "-h")
	utils.Log.Info(output)
}
