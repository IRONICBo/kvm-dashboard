package services

import (
	"fmt"
	"kvm-dashboard/model"
	"kvm-dashboard/utils"
	"kvm-dashboard/vm/agent"
	"strconv"
	"strings"

	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
	"libvirt.org/libvirt-go"
)

// todo: use ssh to get host info
func (svc *Service) GetHostInfo(uuid string) *model.HostInfo {
	hostInfo := &model.HostInfo{}

	username := svc.GetMachineInfo(uuid).Username
	password := svc.GetMachineInfo(uuid).Password
	sshPort := svc.GetMachineInfo(uuid).SshPort
	port, _ := strconv.Atoi(sshPort)
	ip := svc.GetMachineInfo(uuid).Ip

	conn, err := goph.NewConn(
		&goph.Config{
			User:     username,
			Addr:     ip,
			Port:     uint(port),
			Auth:     goph.Password(password),
			Callback: ssh.InsecureIgnoreHostKey(), // set unknown host key callback
		})
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not connect to host: %#v", ip), err)
		return hostInfo
	}

	// get host name
	hostnameOutput, err := utils.GetCommandOutput(conn, "hostname")
	if err != nil {
		utils.Log.Error("Error executing hostname command:", err)
	}
	hostInfo.Hostname = strings.TrimSuffix(string(hostnameOutput), "\n")

	// get os type
	unameOutput, err := utils.GetCommandOutput(conn, "uname", "-a")
	if err != nil {
		utils.Log.Error("Error executing uname command:", err)
	}
	hostInfo.OsType = strings.Split(string(unameOutput), " ")[2]

	// get cpu num
	cpuNumOutput, err := utils.GetCommandOutput(conn, "cat", "/proc/cpuinfo")
	if err != nil {
		utils.Log.Error("Error executing cat /proc/cpuinfo command:", err)
	}
	hostInfo.CPU = strings.Count(string(cpuNumOutput), "processor")

	// get cpu name
	cpuName := strings.Split(string(cpuNumOutput), "\n")[1]
	hostInfo.CPUName = strings.TrimSpace(strings.Split(cpuName, ":")[1])

	// get max mem
	memOutput, err := utils.GetCommandOutput(conn, "cat", "/proc/meminfo")
	if err != nil {
		utils.Log.Error("Error executing cat /proc/meminfo command:", err)
	}
	memTotal := strings.Split(string(memOutput), "\n")[0]
	// utils.Log.Info("memTotal:", strings.TrimSpace(strings.Split(memTotal, ":")[1]))
	mem, err := strconv.Atoi(strings.TrimSpace(strings.Split(strings.TrimSpace(strings.Split(memTotal, ":")[1]), " ")[0]))
	if err != nil {
		utils.Log.Error("Error converting mem to int:", err)
	}
	hostInfo.MaxMem = strconv.Itoa(mem/1024/1024) + "GB"

	return hostInfo
}

func (svc *Service) GetVMInfo(uuid string) *model.VMInfo {
	vmInfo := &model.VMInfo{}

	libvirtURL := svc.GetMachineInfo(uuid).LibvirtUrl
	libvirtAgent, err := agent.NewLibvirtAgent(libvirtURL)
	if err != nil {
		utils.Log.Error("Can not create libvirt agent", err)
		return vmInfo
	}

	info, err := libvirtAgent.GetVMInfo(uuid)
	if err != nil {
		utils.Log.Error("Can not get vm info", err)
		return vmInfo
	}

	vmInfo.Id = info.Id
	vmInfo.Name = info.Name
	vmInfo.UUID = info.UUID
	vmInfo.OsType = info.OsType
	vmInfo.CPU = info.CPU
	vmInfo.IsPersistent = info.IsPersistent
	vmInfo.AutoStart = info.AutoStart
	vmInfo.IpAddress = info.IpAddress

	maxMem := strconv.Itoa(int(info.MaxMem/1024/1024)) + "GB"
	vmInfo.MaxMem = maxMem

	switch info.State {
	case int(libvirt.DOMAIN_RUNNING):
		vmInfo.State = "Running"
	case int(libvirt.DOMAIN_BLOCKED):
		vmInfo.State = "Blocked"
	case int(libvirt.DOMAIN_PAUSED):
		vmInfo.State = "Paused"
	case int(libvirt.DOMAIN_SHUTDOWN):
		vmInfo.State = "Shutdown"
	case int(libvirt.DOMAIN_SHUTOFF):
		vmInfo.State = "Shutoff"
	case int(libvirt.DOMAIN_CRASHED):
		vmInfo.State = "Crashed"
	case int(libvirt.DOMAIN_PMSUSPENDED):
		vmInfo.State = "PMSuspended"
	default:
		vmInfo.State = "No state"
	}

	return vmInfo
}
