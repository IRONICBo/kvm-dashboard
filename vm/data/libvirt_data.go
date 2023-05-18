package data

import "libvirt.org/libvirt-go"

// VM Info
type VMInfo struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	UUID         string `json:"uuid"`
	OsType       string `json:"os_type"`
	State        int    `json:"state"`
	CPU          int    `json:"cpu"`
	MaxMem       uint64 `json:"max_mem"`
	IsPersistent bool   `json:"is_persistent"`
	AutoStart    bool   `json:"auto_start"`
}

// VM collected data
type LibvirtData struct {
	CPUStats       *CPUStats       `json:"cpu_stats"`
	MemoryStats    *MemoryStats    `json:"memory_stats"`
	BlockStats     *BlockStats     `json:"block_stats"`
	InterfaceStats *InterfaceStats `json:"interface_stats"`
}

func NewLibvirtData(domainCPUStats libvirt.DomainCPUStats,
	domainMemoryStats []libvirt.DomainMemoryStat,
	domainBlockStats libvirt.DomainBlockStats,
	domainInterfaceStats libvirt.DomainInterfaceStats) *LibvirtData {
	return &LibvirtData{
		CPUStats:       NewCPUStats(domainCPUStats),
		MemoryStats:    NewMemoryStats(domainMemoryStats),
		BlockStats:     NewBlockStats(domainBlockStats),
		InterfaceStats: NewInterfaceStats(domainInterfaceStats),
	}
}

type CPUStats struct {
	CpuTime    uint64 `json:"cpu_time"`
	UserTime   uint64 `json:"user_time"`
	SystemTime uint64 `json:"system_time"`
	VcpuTime   uint64 `json:"vcpu_time"`
}

func NewCPUStats(domainCPUStats libvirt.DomainCPUStats) *CPUStats {
	return &CPUStats{
		CpuTime:    domainCPUStats.CpuTime,
		UserTime:   domainCPUStats.UserTime,
		SystemTime: domainCPUStats.SystemTime,
		VcpuTime:   domainCPUStats.VcpuTime,
	}
}

type MemoryStats struct {
	SwapIn         uint64 `json:"swap_in"`
	SwapOut        uint64 `json:"swap_out"`
	MajorFault     uint64 `json:"major_fault"`
	MinorFault     uint64 `json:"minor_fault"`
	Unused         uint64 `json:"unused"`
	Available      uint64 `json:"available"`
	ActualBallon   uint64 `json:"actual_ballon"`
	Rss            uint64 `json:"rss"`
	Usable         uint64 `json:"usable"`
	LastUpdate     uint64 `json:"last_update"`
	DiskCaches     uint64 `json:"disk_caches"`
	HugeTLBPgAlloc uint64 `json:"huge_tlb_pg_alloc"`
	HugeTLBPgFail  uint64 `json:"huge_tlb_pg_fail"`
}

func NewMemoryStats(domainMemoryStats []libvirt.DomainMemoryStat) *MemoryStats {
	memoryStats := &MemoryStats{}
	for _, stat := range domainMemoryStats {
		switch libvirt.DomainMemoryStatTags(stat.Tag) {
		case libvirt.DOMAIN_MEMORY_STAT_SWAP_IN:
			memoryStats.SwapIn = stat.Val
		case libvirt.DOMAIN_MEMORY_STAT_SWAP_OUT:
			memoryStats.SwapOut = stat.Val
		case libvirt.DOMAIN_MEMORY_STAT_MAJOR_FAULT:
			memoryStats.MajorFault = stat.Val
		case libvirt.DOMAIN_MEMORY_STAT_MINOR_FAULT:
			memoryStats.MinorFault = stat.Val
		case libvirt.DOMAIN_MEMORY_STAT_UNUSED:
			memoryStats.Unused = stat.Val
		case libvirt.DOMAIN_MEMORY_STAT_AVAILABLE:
			memoryStats.Available = stat.Val
		case libvirt.DOMAIN_MEMORY_STAT_ACTUAL_BALLOON:
			memoryStats.ActualBallon = stat.Val
		case libvirt.DOMAIN_MEMORY_STAT_RSS:
			memoryStats.Rss = stat.Val
		case libvirt.DOMAIN_MEMORY_STAT_USABLE:
			memoryStats.Usable = stat.Val
		case libvirt.DOMAIN_MEMORY_STAT_LAST_UPDATE:
			memoryStats.LastUpdate = stat.Val
		case libvirt.DOMAIN_MEMORY_STAT_DISK_CACHES:
			memoryStats.DiskCaches = stat.Val
		case libvirt.DOMAIN_MEMORY_STAT_HUGETLB_PGALLOC:
			memoryStats.HugeTLBPgAlloc = stat.Val
		case libvirt.DOMAIN_MEMORY_STAT_HUGETLB_PGFAIL:
			memoryStats.HugeTLBPgFail = stat.Val
		}
	}
	return memoryStats
}

type BlockStats struct {
	RdBytes         int64 `json:"rd_bytes"`
	RdReq           int64 `json:"rd_req"`
	RdTotalTimes    int64 `json:"rd_total_times"`
	WrBytes         int64 `json:"wr_bytes"`
	WrReq           int64 `json:"wr_req"`
	WrTotalTimes    int64 `json:"wr_total_times"`
	FlushReq        int64 `json:"flush_req"`
	FlushTotalTimes int64 `json:"flush_total_times"`
	Errs            int64 `json:"errs"`
}

func NewBlockStats(domainBlockStats libvirt.DomainBlockStats) *BlockStats {
	return &BlockStats{
		RdBytes:         domainBlockStats.RdBytes,
		RdReq:           domainBlockStats.RdReq,
		RdTotalTimes:    domainBlockStats.RdTotalTimes,
		WrBytes:         domainBlockStats.WrBytes,
		WrReq:           domainBlockStats.WrReq,
		WrTotalTimes:    domainBlockStats.WrTotalTimes,
		FlushReq:        domainBlockStats.FlushReq,
		FlushTotalTimes: domainBlockStats.FlushTotalTimes,
		Errs:            domainBlockStats.Errs,
	}
}

type InterfaceStats struct {
	RxBytes   int64 `json:"rx_bytes"`
	RxPackets int64 `json:"rx_packets"`
	RxErrs    int64 `json:"rx_errs"`
	RxDrop    int64 `json:"rx_drop"`
	TxBytes   int64 `json:"tx_bytes"`
	TxPackets int64 `json:"tx_packets"`
	TxErrs    int64 `json:"tx_errs"`
	TxDrop    int64 `json:"tx_drop"`
}

func NewInterfaceStats(domainInterfaceStats libvirt.DomainInterfaceStats) *InterfaceStats {
	return &InterfaceStats{
		RxBytes:   domainInterfaceStats.RxBytes,
		RxPackets: domainInterfaceStats.RxPackets,
		RxErrs:    domainInterfaceStats.RxErrs,
		RxDrop:    domainInterfaceStats.RxDrop,
		TxBytes:   domainInterfaceStats.TxBytes,
		TxPackets: domainInterfaceStats.TxPackets,
		TxErrs:    domainInterfaceStats.TxErrs,
		TxDrop:    domainInterfaceStats.TxDrop,
	}
}
