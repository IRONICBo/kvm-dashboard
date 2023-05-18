package utils

import (
	"fmt"
	"kvm-dashboard/vm/data"
	"testing"

	"github.com/magiconair/properties/assert"
	"libvirt.org/libvirt-go"
)

func TestStructToMap(t *testing.T) {
	data := data.NewLibvirtData(libvirt.DomainCPUStats{}, []libvirt.DomainMemoryStat{}, libvirt.DomainBlockStats{}, libvirt.DomainInterfaceStats{})
	res, err := StructToMapWithJSONTag(data)

	Log.Info(fmt.Sprintf("%#v", res))
	assert.Equal(t, err, nil)

	// time="2023-05-16 17:20:04.408" level=info msg="map[string]interface {}{\"block_stats\":map[string]interface {}{\"errs\":0, \"flush_req\":0, \"flush_total_times\":0, \"rd_bytes\":0, \"rd_req\":0, \"rd_total_times\":0, \"wr_bytes\":0, \"wr_req\":0, \"wr_total_times\":0}, \"cpu_stats\":map[string]interface {}{\"cpu_time\":0, \"system_time\":0, \"user_time\":0, \"vcpu_time\":0}, \"interface_stats\":map[string]interface {}{\"rx_bytes\":0, \"rx_drop\":0, \"rx_errs\":0, \"rx_packets\":0, \"tx_bytes\":0, \"tx_drop\":0, \"tx_errs\":0, \"tx_packets\":0}, \"memory_stats\":map[string]interface {}{\"actual_ballon\":0, \"available\":0, \"disk_caches\":0, \"huge_tlb_pg_alloc\":0, \"huge_tlb_pg_fail\":0, \"last_update\":0, \"major_fault\":0, \"minor_fault\":0, \"rss\":0, \"swap_in\":0, \"swap_out\":0, \"unused\":0, \"usable\":0}}" file=" convert_test.go:17"
}

func TestFlattenMap(t *testing.T) {
	data := data.NewLibvirtData(libvirt.DomainCPUStats{}, []libvirt.DomainMemoryStat{}, libvirt.DomainBlockStats{}, libvirt.DomainInterfaceStats{})
	res, err := StructToMapWithJSONTag(data)
	Log.Info(fmt.Sprintf("%#v", res))
	assert.Equal(t, err, nil)

	// flatten
	flatten := FlattenMap(res, "+")
	Log.Info(fmt.Sprintf("%#v", flatten))

	// convert_test.go:30 map[string]interface {}{"block_stats+errs":0, "block_stats+flush_req":0, "block_stats+flush_total_times":0, "block_stats+rd_bytes":0, "block_stats+rd_req":0, "block_stats+rd_total_times":0, "block_stats+wr_bytes":0, "block_stats+wr_req":0, "block_stats+wr_total_times":0, "cpu_stats+cpu_time":0, "cpu_stats+system_time":0, "cpu_stats+user_time":0, "cpu_stats+vcpu_time":0, "interface_stats+rx_bytes":0, "interface_stats+rx_drop":0, "interface_stats+rx_errs":0, "interface_stats+rx_packets":0, "interface_stats+tx_bytes":0, "interface_stats+tx_drop":0, "interface_stats+tx_errs":0, "interface_stats+tx_packets":0, "memory_stats+actual_ballon":0, "memory_stats+available":0, "memory_stats+disk_caches":0, "memory_stats+huge_tlb_pg_alloc":0, "memory_stats+huge_tlb_pg_fail":0, "memory_stats+last_update":0, "memory_stats+major_fault":0, "memory_stats+minor_fault":0, "memory_stats+rss":0, "memory_stats+swap_in":0, "memory_stats+swap_out":0, "memory_stats+unused":0, "memory_stats+usable":0}
}
