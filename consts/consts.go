package consts

const (
	LIBVIRT_URL    = "qemu:///system"
	INTERFACE_NAME = "enp1s0"
)

const (
	VM_DATA_REALTIME    = "vm_data_realtime"    // realtime data
	VM_DATA_NEARLY_HOUR = "vm_data_nearly_hour" // nearly one hour data
	VM_DATA_NEARLY_DAY  = "vm_data_nearly_day"  // nearly one day data

	VISION_COUNT                       = 20                                                          // 20 point for echarts
	VM_DATA_REALTIME_INTERVAL_INTERVAL = 3                                                           // 3 second
	VM_DATA_NEARLY_5M_INTERVAL         = 5 * 60 / VISION_COUNT / VM_DATA_REALTIME_INTERVAL_INTERVAL  // 5 second
	VM_DATA_NEARLY_1H_INTERVAL         = 60 * 60 / VISION_COUNT / VM_DATA_REALTIME_INTERVAL_INTERVAL // 60 second
)

const (
	PERIOD_1M = "1m"
	PERIOD_5M = "5m"
	PERIOD_1H = "1h"
)

const (
	AGG_MEAN   = "mean"
	AGG_MAX    = "max"
	AGG_MIN    = "min"
	AGG_MEDIAN = "median"
)

// Unit
var (
	// todo: add index unit
	UNIT = map[string]string{}
)
