package vm_utils

import (
	"kvm-dashboard/vm/data"
	"strings"
)

func ParseDstatData(out string) string {
	// todo: parse dstat data

	return ""
}

func ParseTopData(out string) []*data.ProcessData {
	processes := make([]*data.ProcessData, 0)
	strs := strings.Split(out, "\n")

	// top - 23:39:18 up 6 days, 12:10,  1 user,  load average: 0.00, 0.00, 0.00
	// Tasks: 127 total,   1 running, 126 sleeping,   0 stopped,   0 zombie
	// %Cpu(s):  0.0 us,  6.2 sy,  0.0 ni, 93.8 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st
	// MiB Mem :    988.0 total,    114.9 free,    445.2 used,    427.9 buff/cache
	// MiB Swap:   2047.9 total,   1979.5 free,     68.4 used.    324.1 avail Mem
	//
	// PID USER      PR  NI    VIRT    RES    SHR S  %CPU  %MEM     TIME+ COMMAND
	for index, line := range strs[7:] {
		fields := strings.Fields(line)
		if len(fields) >= 12 {
			process := data.NewProcessData(
				fields[0],
				fields[1],
				fields[2],
				fields[3],
				fields[4],
				fields[5],
				fields[6],
				fields[7],
				fields[8],
				fields[9],
				fields[10],
				fields[11],
			)
			processes = append(processes, process)
		}

		// index start from 0 => 7
		if index >= 19 {
			break
		}
	}

	return processes
}
