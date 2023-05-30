package consts

import "kvm-dashboard/model"

const (
	// LIBVIRT_URL    = "qemu:///system"
	LIBVIRT_URL    = "qemu+tcp://localhost:16509/system" // default url
	INTERFACE_NAME = "enp1s0"
	USERNAME       = "root"
	PASSWORD       = "123123"
	PORT           = 22
)

const (
	VM_DATA_REALTIME    = "vm_data_realtime"    // realtime data
	VM_DATA_NEARLY_HOUR = "vm_data_nearly_hour" // nearly one hour data
	VM_DATA_NEARLY_DAY  = "vm_data_nearly_day"  // nearly one day data

	VISION_COUNT               = 20                                                 // 20 point for echarts
	VM_DATA_REALTIME_INTERVAL  = 3                                                  // 3 second
	VM_DATA_NEARLY_5M_INTERVAL = 5 * 60 / VISION_COUNT / VM_DATA_REALTIME_INTERVAL  // 5 second
	VM_DATA_NEARLY_1H_INTERVAL = 60 * 60 / VISION_COUNT / VM_DATA_REALTIME_INTERVAL // 60 second
)

const (
	PERIOD_1M  = "1m"
	PERIOD_5M  = "5m"
	PERIOD_1H  = "1h"
	PERIOD_30D = "30d"
)

const (
	AGG_MEAN   = "mean"
	AGG_MAX    = "max"
	AGG_MIN    = "min"
	AGG_MEDIAN = "median"
)

const (
	// measurement
	VM_MEASUREMENT     = "vm"
	SIMPLE_MEASUREMENT = "simple"
	ALERT_MEASUREMENT  = "alert"
)

// Unit
var (
	// todo: add index unit
	UNIT = map[string]string{}
)

var (
	VM_THRESHOLD         = map[string]*model.Threshold{}
	CPU_USAGE_THRESHOLD  = 80
	MEM_USAGE_THRESHOLD  = 80
	DISK_USAGE_THRESHOLD = 80
)

// Realtime data
var ()
