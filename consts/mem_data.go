package consts

import "kvm-dashboard/model"

// Unit
var (
	// todo: add index unit
	UNIT = map[string]string{}
)

// threshold
var (
	VM_THRESHOLD         = map[string]*model.Threshold{}
	CPU_USAGE_THRESHOLD  = 80
	MEM_USAGE_THRESHOLD  = 80
	DISK_USAGE_THRESHOLD = 80
)

// machine info
var (
	// uuid, host or vm
	MACHINE_MAP = map[string]*model.Machine{}
)

// Realtime data
var ()
