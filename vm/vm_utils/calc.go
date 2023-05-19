package vm_utils

import (
	"errors"
	"fmt"
	"kvm-dashboard/utils"
	"kvm-dashboard/vm/data"
	"strings"
)

func ParseCPULoad(prev, current string) (float64, error) {
	prevStats, err := parseSimpleCPUStats(prev)
	if err != nil {
		utils.Log.Error("Can not parse cpu stats", err)
		return 0, err
	}
	currentStats, err := parseSimpleCPUStats(current)
	if err != nil {
		utils.Log.Error("Can not parse cpu stats", err)
		return 0, err
	}

	totalDiff := getTotalDiff(prevStats, currentStats)
	idleDiff := getIdleDiff(prevStats, currentStats)

	cpuUsage := 100.0 * (1.0 - float64(idleDiff)/float64(totalDiff))

	return cpuUsage, nil
}

func parseSimpleCPUStats(str string) (*data.SimpleCPUStats, error) {
	str = strings.Split(str, "\n")[0]

	stats := &data.SimpleCPUStats{}
	var cpu string

	// format:
	// cpu xxx xxx xxx

	// _, err := fmt.Sscanf(data, "%d %d %d %d %d %d %d %d %d %d",
	// 	&stats.User, &stats.Nice, &stats.System, &stats.Idle, &stats.Iowait,
	// 	&stats.Irq, &stats.Softirq, &stats.Steal, &stats.Guest, &stats.GuestNice)
	_, err := fmt.Sscanf(str, "%s %d %d %d %d", &cpu, &stats.User, &stats.Nice, &stats.System, &stats.Idle)
	if err != nil {
		utils.Log.Error("Can not parse cpu stats", err)
		return nil, err
	}

	return stats, nil
}

func getTotalDiff(prevStats, currStats *data.SimpleCPUStats) uint64 {
	prevTotal := prevStats.User + prevStats.Nice + prevStats.System + prevStats.Idle
	currTotal := currStats.User + currStats.Nice + currStats.System + currStats.Idle
	return currTotal - prevTotal
}

func getIdleDiff(prevStats, currStats *data.SimpleCPUStats) uint64 {
	prevIdle := prevStats.Idle
	currIdle := currStats.Idle
	return currIdle - prevIdle
}

func ParseMemLoad(str string) (float64, error) {
	lines := strings.Split(str, "\n")
	if len(lines) < 3 {
		utils.Log.Error("Can not parse mem stats", errors.New("invalid mem stats"))
		return 0, errors.New("invalid mem stats")
	}

	fields := strings.Fields(lines[1]) // use whitespace as delimiter
	if len(fields) < 7 {
		utils.Log.Error("Can not parse mem stats", errors.New("invalid mem stats"))
		return 0, errors.New("invalid mem stats")
	}

	var memStats data.SimpleMemoryStats
	memStats.Total, _ = parseMemorySize(fields[1])
	memStats.Used, _ = parseMemorySize(fields[2])
	memStats.Free, _ = parseMemorySize(fields[3])
	memStats.Shared, _ = parseMemorySize(fields[4])
	memStats.Buffer, _ = parseMemorySize(fields[5])
	memStats.Available, _ = parseMemorySize(fields[6])

	// usage
	memUsage := 100.0 * (1.0 - memStats.Available/memStats.Total)
	return memUsage, nil
}

func parseMemorySize(size string) (float64, error) {
	size = strings.TrimSpace(size)

	var value float64
	var unit string
	_, err := fmt.Sscanf(size, "%f%s", &value, &unit)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func ParseDiskLoad(str string) (float64, error) {
	lines := strings.Split(str, "\n")
	if len(lines) < 2 {
		utils.Log.Error("Can not parse disk stats", errors.New("invalid disk stats"))
		return 0, errors.New("invalid disk stats")
	}

	var diskStats []*data.SimpleDiskStats
	for _, line := range lines[1:] {
		fields := strings.Fields(line)
		if len(fields) < 6 {
			// utils.Log.Error("Can not parse disk stats", errors.New("invalid disk stats"))
			// return 0, errors.New("invalid disk stats")
			continue
		}

		size, _ := parseDiskSize(fields[1])
		used, _ := parseDiskSize(fields[2])
		available, _ := parseDiskSize(fields[3])

		diskStat := &data.SimpleDiskStats{
			Filesystem: fields[0],
			Size:       size,
			Used:       used,
			Available:  available,
			Usage:      fields[4],
			MountedOn:  fields[5],
		}
		diskStats = append(diskStats, diskStat)
	}

	return calcTotalUsage(diskStats), nil
}

func calcTotalUsage(diskStats []*data.SimpleDiskStats) float64 {
	var totalUsage, totalSize float64
	for _, diskStat := range diskStats {
		if diskStat.Filesystem == "tmpfs" || diskStat.Filesystem == "devtmpfs" {
			continue
		}

		totalUsage += diskStat.Used
		totalSize += diskStat.Size
	}

	return 100.0 * (totalUsage / totalSize)
}

func parseDiskSize(size string) (float64, error) {
	size = strings.TrimSpace(size)

	var value float64
	var unit string
	_, err := fmt.Sscanf(size, "%f%s", &value, &unit)
	if err != nil {
		return 0, err
	}

	// base on K
	switch unit {
	case "K":
		value *= 1
	case "M":
		value *= 1 * 1024
	case "G":
		value *= 1 * 1024 * 1024
	case "T":
		value *= 1 * 1024 * 1024 * 1024
	}

	return value, nil
}

func ParseNetLoad(prev, current, bandwith string, interval int) (float64, float64, error) {
	prevRx, prevTx, err := parseNetstat(prev)
	// utils.Log.Info(fmt.Sprintf("prevRx: %#v prevTx: %#v", prevRx, prevTx))
	if err != nil {
		utils.Log.Error("Can not parse net stats", err)
		return 0, 0, err
	}
	currRx, currTx, err := parseNetstat(current)
	// utils.Log.Info(fmt.Sprintf("currRx: %#v currTx: %#v", currRx, currTx))
	if err != nil {
		utils.Log.Error("Can not parse net stats", err)
		return 0, 0, err
	}

	// bandwidth, err := parseBandwith(bandwith)
	// if err != nil {
	// 	utils.Log.Error("Can not parse net stats", err)
	// 	return 0, 0, err
	// }

	rxRate := (float64(currRx-prevRx) / float64(interval))
	txRate := (float64(currTx-prevTx) / float64(interval))

	return rxRate, txRate, nil
}

func parseNetstat(str string) (uint64, uint64, error) {
	// 解析命令输出
	lines := strings.Split(str, "\n")
	rxLine := lines[3]
	txLine := lines[5]

	var rxBytes, txBytes uint64
	_, err := fmt.Sscanf(rxLine, "%d", &rxBytes)
	if err != nil {
		return 0, 0, err
	}
	_, err = fmt.Sscanf(txLine, "%d", &txBytes)
	if err != nil {
		return 0, 0, err
	}

	return rxBytes, txBytes, nil
}

func parseBandwith(bandwith string) (uint64, error) {
	var bandwidth uint64
	lines := strings.Split(bandwith, "\n")
	for _, line := range lines {
		if strings.Contains(line, "Speed:") {
			fields := strings.Fields(line)
			if len(fields) >= 2 {
				speed := fields[1]
				speed = strings.TrimSpace(speed)

				if speed == "Unknown!" {
					return 0, nil
				}
				_, err := fmt.Sscanf(speed, "%d", &bandwidth)
				if err != nil {
					return 0, err
				}
			}
		}
	}

	// convert to byte
	bandwidth = bandwidth * 1024 * 1024 / 8 // convert to byte

	return 0, errors.New("can not find bandwidth")
}
