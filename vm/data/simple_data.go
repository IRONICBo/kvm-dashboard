package data

type SimpleData struct {
	CPUUsage  float64 `json:"cpu_usage"`
	MemUsage  float64 `json:"mem_usage"`
	DiskUsage float64 `json:"disk_usage"`
	NetRxRate float64 `json:"net_rx_rate"`
	NetTxRate float64 `json:"net_tx_rate"`
}

func NewSimpleData(cpuUsage, memUsage, diskUsage, netRxRate, netTxRate float64) *SimpleData {
	return &SimpleData{
		CPUUsage:  cpuUsage,
		MemUsage:  memUsage,
		DiskUsage: diskUsage,
		NetRxRate: netRxRate,
		NetTxRate: netTxRate,
	}
}

type SimpleCPUStats struct {
	User   uint64
	Nice   uint64
	System uint64
	Idle   uint64
}

type SimpleMemoryStats struct {
	Total     float64
	Used      float64
	Free      float64
	Shared    float64
	Buffer    float64
	Available float64
}

type SimpleDiskStats struct {
	Filesystem string
	Size       float64
	Used       float64
	Available  float64
	Usage      string
	MountedOn  string
}

type SimpleNetStats struct {
	RXDiff uint64
	TXDiff uint64
}
