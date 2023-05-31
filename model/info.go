package model

type VMInfo struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	UUID         string `json:"uuid"`
	OsType       string `json:"os_type"`
	State        string `json:"state"`
	CPU          int    `json:"cpu"`
	MaxMem       string `json:"max_mem"`
	IsPersistent bool   `json:"is_persistent"`
	AutoStart    bool   `json:"auto_start"`
	IpAddress    string `json:"ip_address"`
}

type HostInfo struct {
	Name       string `json:"name"`
	Hostname   string `json:"hostname"`
	UUID       string `json:"uuid"`
	Ip         string `json:"ip"`
	OsType     string `json:"os_type"`
	CPU        int    `json:"cpu"`
	CPUName    string `json:"cpu_name"`
	MaxMem     string `json:"max_mem"`
	SshPort    string `json:"ssh_port"`
	LibvirtUrl string `json:"libvirt_url"`
}
