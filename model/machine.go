package model

type Machine struct {
	Name       string `json:"name"`
	Uuid       string `json:"uuid"`
	Ip         string `json:"ip"`
	SshPort    string `json:"ssh_port"`
	LibvirtUrl string `json:"libvirt_url"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

func NewMachine(name, uuid, ip, sshPort, libvirtUrl, username, password string) *Machine {
	return &Machine{
		Name:       name,
		Uuid:       uuid,
		Ip:         ip,
		SshPort:    sshPort,
		LibvirtUrl: libvirtUrl,
		Username:   username,
		Password:   password,
	}
}

type MachineInfo struct {
	HostInfo []*Machine
	VmInfo   []*Machine
}

func NewMachineInfo(hostInfo []*Machine, vmInfo []*Machine) *MachineInfo {
	return &MachineInfo{
		HostInfo: hostInfo,
		VmInfo:   vmInfo,
	}
}

// Threshold
type Threshold struct {
	CpuUsageThreshold  int `json:"cpu_usage_threshold"`
	MemUsageThreshold  int `json:"mem_usage_threshold"`
	DiskUsageThreshold int `json:"disk_usage_threshold"`
}

func NewThreshold(cpuUsageThreshold, memUsageThreshold, diskUsageThreshold int) *Threshold {
	return &Threshold{
		CpuUsageThreshold:  cpuUsageThreshold,
		MemUsageThreshold:  memUsageThreshold,
		DiskUsageThreshold: diskUsageThreshold,
	}
}
