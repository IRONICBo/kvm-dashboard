package agent

import (
	"fmt"
	"kvm-dashboard/utils"
	"kvm-dashboard/vm/data"

	libvirt "libvirt.org/libvirt-go"
)

type LibvirtAgent struct {
	*AgentInfo
	AgentInterface
	URL string // Libvirt Controller URL
	// 这里改成conn应该会比较好，可以在Stop之类的方法里面获取到

	LibvirtData chan *data.LibvirtData
}

func NewLibvirtAgent(agentInfo *AgentInfo) *LibvirtAgent {
	return &LibvirtAgent{
		AgentInfo:   agentInfo,
		LibvirtData: make(chan *data.LibvirtData, 100), // set the buffer size to 100
	}
}

func (la *LibvirtAgent) Start() {
	conn, err := libvirt.NewConnect(la.URL)
	if err != nil {
		utils.LogWithError(fmt.Sprintf("Can not connect to libvirt: %#v", la.AgentInfo), err)
		return
	}
	// defer conn.Close()

	doms, _ := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
	for _, do := range doms {
		utils.LogWithInfo(do.GetName())
	}

	dom, err := conn.LookupDomainByUUIDString(la.UUID)
	if err != nil {
		utils.LogWithError(fmt.Sprintf("Can not get dom: %#v", la.AgentInfo), err)
		return
	}

	// Ref: https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetCPUStats
	// flags set 0 => unused
	cpu_stat, err := dom.GetCPUStats(-1, 1, 0)
	if err != nil {
		utils.LogWithInfo(err)
		return
	}

	// 有问题。。。
	utils.LogWithInfo(cpu_stat)
}

func (la *LibvirtAgent) Stop() {
	// conn.Close()
}

func (la *LibvirtAgent) Restart() {
}

func (la *LibvirtAgent) Status() {
}
