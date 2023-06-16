package vm_utils

import (
	"fmt"
	"kvm-dashboard/consts"
	"kvm-dashboard/utils"

	"github.com/melbahja/goph"
)

// Add arp entry to host machine
// vm use `ping host_ip` to get host mac connection
// and host's arp table contains vm's mac address
func AddArpEntry(hostClient *goph.Client, uuid string) error {
	// get host local ip
	vmIp, err := GetCommandOutput(hostClient, "ifconfig | grep 'inet 192' | awk 'NR==1 {print $2}'")
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Get host ip failed: %v", err))
		return err
	}

	// get virsh console
	err = RunCommand(hostClient, "sudo virsh console --force ", uuid)
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Get host ip failed: %v", err))
		return err
	}

	// send username and password
	err = RunCommand(hostClient, consts.USERNAME)
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Set USERNAME failed: %v", err))
		return err
	}
	err = RunCommand(hostClient, consts.PASSWORD)
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Set PASSWORD failed: %v", err))
		return err
	}

	// ping
	err = RunCommand(hostClient, "ping", vmIp)
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Get host ip failed: %v", err))
		return err
	}

	return nil
}

func InitVmNetwork() {

}
