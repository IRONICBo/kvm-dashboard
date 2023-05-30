package model

type MachineInfo struct {
	HostInfo []*Host
	VmInfo   []*Vm
}

func NewMachineInfo(hostInfo []*Host, vmInfo []*Vm) *MachineInfo {
	return &MachineInfo{
		HostInfo: hostInfo,
		VmInfo:   vmInfo,
	}
}

type Host struct {
	Name       string `json:"name"`
	Uuid       string `json:"uuid"`
	Ip         string `json:"ip"`
	SshPort    string `json:"ssh_port"`
	LibvirtUrl string `json:"libvirt_url"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

func NewHost(name, uuid, ip, sshPort, libvirtUrl, username, password string) *Host {
	return &Host{
		Name:       name,
		Uuid:       uuid,
		Ip:         ip,
		SshPort:    sshPort,
		LibvirtUrl: libvirtUrl,
		Username:   username,
		Password:   password,
	}
}

type Vm struct {
	Name       string `json:"name"`
	Uuid       string `json:"uuid"`
	Ip         string `json:"ip"`
	SshPort    string `json:"ssh_port"`
	LibvirtUrl string `json:"libvirt_url"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

func NewVm(name, uuid, ip, sshPort, libvirtUrl, username, password string) *Vm {
	return &Vm{
		Name:       name,
		Uuid:       uuid,
		Ip:         ip,
		SshPort:    sshPort,
		LibvirtUrl: libvirtUrl,
		Username:   username,
		Password:   password,
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
