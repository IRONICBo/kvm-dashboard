package agent

import (
	"errors"
	"fmt"
	"kvm-dashboard/consts"
	"kvm-dashboard/utils"
	"kvm-dashboard/vm/data"
	"time"

	libvirt "libvirt.org/libvirt-go"
)

type LibvirtAgent struct {
	// *AgentInfo
	AgentInterface
	URL  string
	conn *libvirt.Connect
	// 这里改成conn应该会比较好，可以在Stop之类的方法里面获取到

	LibvirtData chan *data.LibvirtData
}

func NewLibvirtAgent(url string) (*LibvirtAgent, error) {
	agent := &LibvirtAgent{
		// LibvirtData: make(chan *data.LibvirtData, 100), // set the buffer size to 100
		URL: url,
	}

	// connect to vm
	conn, err := libvirt.NewConnect(url)
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not connect to libvirt: %#v", url), err)
		return nil, err
	}
	agent.conn = conn

	return agent, nil
}

// Init VM utils
// todo: use console to setup vm network and install essential packages
func (la *LibvirtAgent) VmSetUp(uuid, username, password string) error {
	// get interface name
	dom, err := la.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not get dom: %#v", uuid), err)
		return err
	}
	defer dom.Free()

	interfaces, err := dom.ListAllInterfaceAddresses(libvirt.DOMAIN_INTERFACE_ADDRESSES_SRC_AGENT)
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not get interfaces: %#v", uuid), err)
		return err
	}
	if len(interfaces) <= 1 {
		utils.Log.Error(fmt.Sprintf("Can not get enough interfaces: %#v", uuid), err)
		return errors.New("can not get enough interfaces")
	}
	interface_name := interfaces[1].Name
	utils.Log.Info(fmt.Sprintf("interface_name: %#v", interfaces))
	utils.Log.Info(fmt.Sprintf("interface_name: %#v", interface_name))

	stream, err := la.conn.NewStream(libvirt.STREAM_NONBLOCK)
	if err != nil {
		fmt.Println("Failed to create stream:", err)
		return err
	}

	// la.conn.NewStream()
	err = dom.OpenConsole("", stream, libvirt.DOMAIN_CONSOLE_FORCE)
	if err != nil {
		fmt.Println("Failed to open console:", err)
		return err
	}

	_, err = stream.Send([]byte("\n"))
	if err != nil {
		fmt.Println("Failed to send login command:", err)
		return err
	}
	// get result
	buf := make([]byte, 1024)
	_, err = stream.Recv(buf)
	if err != nil {
		fmt.Println("Failed to receive output:", err)
		return err
	}

	// login
	loginCommand := fmt.Sprintf("\n%s\n%s\n", username, password)
	_, err = stream.Send([]byte(loginCommand))
	if err != nil {
		fmt.Println("Failed to send login command:", err)
		return err
	}
	// get result
	buf = make([]byte, 1024)
	_, err = stream.Recv(buf)
	if err != nil {
		fmt.Println("Failed to receive output:", err)
		return err
	}

	// connect to vm
	command := fmt.Sprintf("nmcli connection up %s\n", interface_name)
	_, err = stream.Send([]byte(command))
	if err != nil {
		fmt.Println("Failed to send command:", err)
		return err
	}

	// get result
	buf = make([]byte, 1024)
	_, err = stream.Recv(buf)
	if err != nil {
		fmt.Println("Failed to receive output:", err)
		return err
	}

	return nil
}

// collect data
func (la *LibvirtAgent) Start(uuid string) {
	la.LibvirtData = make(chan *data.LibvirtData, 100) // set the buffer size to 100

	go func() {
		defer close(la.LibvirtData) // close libvirtData

		dom, err := la.conn.LookupDomainByUUIDString(uuid)
		if err != nil {
			utils.Log.Error(fmt.Sprintf("Can not get dom: %#v", uuid), err)
			return
		}
		defer dom.Free()

		for {
			// Ref: https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetCPUStats
			// flags set 0 => unused
			cpu_stats, err := dom.GetCPUStats(-1, 1, 0) // DomainCPUStats
			if err != nil {
				utils.Log.Error(fmt.Sprintf("Can not get cpu_stats: %#v", uuid), err)
				return
			}
			memory_stats, err := dom.MemoryStats(uint32(libvirt.DOMAIN_MEMORY_STAT_NR), 0) // MemoryStats
			if err != nil {
				utils.Log.Error(fmt.Sprintf("Can not get memory_stats: %#v", uuid), err)
				return
			}
			block_stats, err := dom.BlockStats("") // set "" to get all block stats
			if err != nil {
				utils.Log.Error(fmt.Sprintf("Can not get block_stats: %#v", uuid), err)
				return
			}

			// get mac address
			ifaces, err := dom.ListAllInterfaceAddresses(libvirt.DOMAIN_INTERFACE_ADDRESSES_SRC_AGENT)
			if err != nil {
				utils.Log.Error(fmt.Sprintf("Can not get ifaces: %#v", uuid), err)
				return
			}
			var MAC_address string
			for _, iface := range ifaces {
				if len(iface.Hwaddr) > 0 && iface.Name == consts.INTERFACE_NAME {
					MAC_address = iface.Hwaddr
				}
			}
			interface_stats, err := dom.InterfaceStats(MAC_address) // use mac address or interface name
			if err != nil {
				utils.Log.Error(fmt.Sprintf("Can not get interface_stats: %#v", uuid), err)
				return
			}

			libvirtData := data.NewLibvirtData(cpu_stats[0], memory_stats, *block_stats, *interface_stats)
			la.LibvirtData <- libvirtData // send data

			// delay
			time.Sleep(time.Duration(consts.VM_DATA_REALTIME_INTERVAL_INTERVAL) * time.Second)
		}
	}()
}

func (la *LibvirtAgent) Stop() {
	la.conn.Close()
}

func (la *LibvirtAgent) Restart() error {
	// reconnect to vm
	conn, err := libvirt.NewConnect(la.URL)
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not connect to libvirt: %#v", la.URL), err)
		return err
	}

	la.conn = conn
	return nil
}

func (la *LibvirtAgent) Status() {
}

func (la *LibvirtAgent) GetVMInfo(uuid string) (*data.VMInfo, error) {
	dom, err := la.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not get dom: %#v", uuid), err)
		return nil, err
	}
	defer dom.Free()

	id, err := dom.GetID()
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not get id: %#v", uuid), err)
	}

	name, err := dom.GetName()
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not get name: %#v", uuid), err)
	}

	osType, err := dom.GetOSType()
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not get osType: %#v", uuid), err)
	}

	state, _, err := dom.GetState()
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not get state: %#v", uuid), err)
	}

	maxMem, err := dom.GetMaxMemory()
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not get maxMem: %#v", uuid), err)
	}

	cpuInfo, err := dom.GetVcpus()
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not get cpuInfo: %#v", uuid), err)
	}
	cpu := len(cpuInfo)

	isPersistent, err := dom.IsPersistent()
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not get isPersistent: %#v", uuid), err)
	}

	autoStart, err := dom.GetAutostart()
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not get autoStart: %#v", uuid), err)
	}

	vmInfo := &data.VMInfo{
		Id:           id,
		Name:         name,
		UUID:         uuid,
		OsType:       osType,
		State:        int(state),
		CPU:          cpu,
		MaxMem:       maxMem,
		IsPersistent: isPersistent,
		AutoStart:    autoStart,
	}

	return vmInfo, err
}

// todo: start vm and start monitoring
// show progress in controller
func (la *LibvirtAgent) StartVM(uuid string) error {
	dom, err := la.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not get dom: %#v", uuid), err)
		return err
	}
	defer dom.Free()

	err = dom.Create()
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not start vm: %#v", uuid), err)
		return err
	}
	return err
}

func (la *LibvirtAgent) StopVM(uuid string) error {
	dom, err := la.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not get dom: %#v", uuid), err)
		return err
	}
	defer dom.Free()

	err = dom.Shutdown()
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not stop vm: %#v", uuid), err)
		return err
	}
	return err
}

func (la *LibvirtAgent) SuspendVM(uuid string) error {
	dom, err := la.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not get dom: %#v", uuid), err)
		return err
	}
	defer dom.Free()

	err = dom.Suspend()
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not suspend vm: %#v", uuid), err)
		return err
	}
	return err
}

func (la *LibvirtAgent) ResumeVM(uuid string) error {
	dom, err := la.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not get dom: %#v", uuid), err)
		return err
	}
	defer dom.Free()

	err = dom.Resume()
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not resume vm: %#v", uuid), err)
		return err
	}
	return err
}
