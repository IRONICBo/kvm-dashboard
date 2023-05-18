package agent

import (
	"context"
	"errors"
	"fmt"
	"kvm-dashboard/consts"
	"kvm-dashboard/utils"
	"kvm-dashboard/vm/data"
	"kvm-dashboard/vm/vm_utils"
	"time"

	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
)

var SimpleAgents map[string]*SimpleAgent

type SimpleAgent struct {
	*AgentInfo
	AgentInterface
	conn *goph.Client

	Ctx        context.Context // one ctx for one goroutine
	CancelFunc context.CancelFunc
	SimpleData chan *data.SimpleData
}

func NewSimpleAgent(agentInfo *AgentInfo) (*SimpleAgent, error) {
	if SimpleAgents == nil {
		SimpleAgents = make(map[string]*SimpleAgent)
	}

	conn, err := goph.NewConn(
		&goph.Config{
			User:     agentInfo.User,
			Addr:     agentInfo.Ip,
			Port:     agentInfo.Port,
			Auth:     goph.Password(agentInfo.Password),
			Callback: ssh.InsecureIgnoreHostKey(), // set unknown host key callback
		})
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not connect to vm: %#v", agentInfo), err)
		return nil, err
	}

	// create context
	ctx, cancelFunc := context.WithCancel(context.Background())

	agent := &SimpleAgent{
		AgentInfo:  agentInfo,
		conn:       conn,
		Ctx:        ctx,
		CancelFunc: cancelFunc,
	}

	return agent, nil
}

func (sa *SimpleAgent) Start(uuid string) {
	// store started agent
	SimpleAgents[uuid] = sa

	// defer sa.conn.Close()
	sa.SimpleData = make(chan *data.SimpleData, 100) // set the buffer size to 100

	cpuPrev, err := vm_utils.GetCommandOutput(sa.conn, "cat", "/proc/stat")
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not get current proc stat result on vm: %#v", sa.AgentInfo), err)
		return
	}
	// net
	netPrev, err := vm_utils.GetCommandOutput(sa.conn, "ip", "-s link show", consts.INTERFACE_NAME)
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not get next proc stat result on vm: %#v", sa.AgentInfo), err)
		return
	}

	go func() {
		defer close(sa.SimpleData)

		for {
			select {
			case <-sa.Ctx.Done():
				utils.Log.Info(fmt.Sprintf("Stop simple agent goroutine: %#v", uuid))
				return
			default:
				// delay for interval to get diff
				time.Sleep(time.Duration(consts.VM_DATA_REALTIME_INTERVAL) * time.Second)

				// cpu
				cpuCurrent, err := vm_utils.GetCommandOutput(sa.conn, "cat", "/proc/stat")
				if err != nil {
					utils.Log.Error(fmt.Sprintf("Can not get next proc stat result on vm: %#v", sa.AgentInfo), err)
					return
				}
				cpuLoad, err := vm_utils.ParseCPULoad(cpuPrev, cpuCurrent)
				if err != nil {
					utils.Log.Error(fmt.Sprintf("Can not parse cpu load on vm: %#v", sa.AgentInfo), err)
					return
				}

				// mem
				mem, err := vm_utils.GetCommandOutput(sa.conn, "free", "-h")
				if err != nil {
					utils.Log.Error(fmt.Sprintf("Can not get next proc stat result on vm: %#v", sa.AgentInfo), err)
					return
				}
				memLoad, err := vm_utils.ParseMemLoad(mem)
				if err != nil {
					utils.Log.Error(fmt.Sprintf("Can not parse mem load on vm: %#v", sa.AgentInfo), err)
					return
				}

				// disk
				disk, err := vm_utils.GetCommandOutput(sa.conn, "df", "-h")
				if err != nil {
					utils.Log.Error(fmt.Sprintf("Can not get next proc stat result on vm: %#v", sa.AgentInfo), err)
					return
				}
				diskLoad, err := vm_utils.ParseDiskLoad(disk)
				if err != nil {
					utils.Log.Error(fmt.Sprintf("Can not parse disk load on vm: %#v", sa.AgentInfo), err)
					return
				}

				// net
				netCurrent, err := vm_utils.GetCommandOutput(sa.conn, "ip", "-s link show", consts.INTERFACE_NAME)
				if err != nil {
					utils.Log.Error(fmt.Sprintf("Can not get next proc stat result on vm: %#v", sa.AgentInfo), err)
					return
				}
				bandwith, err := vm_utils.GetCommandOutput(sa.conn, "ethtool", consts.INTERFACE_NAME)
				if err != nil {
					utils.Log.Error(fmt.Sprintf("Can not get next proc stat result on vm: %#v", sa.AgentInfo), err)
					return
				}
				rxRate, txRate, err := vm_utils.ParseNetLoad(netPrev, netCurrent, bandwith, consts.VM_DATA_REALTIME_INTERVAL)
				if err != nil {
					utils.Log.Error(fmt.Sprintf("Can not parse net rate on vm: %#v", sa.AgentInfo), err)
					return
				}

				// send data
				simpleData := data.NewSimpleData(cpuLoad, memLoad, diskLoad, rxRate, txRate)
				utils.Log.Info(fmt.Sprintf("Simple data: %#v", simpleData))
				sa.SimpleData <- simpleData

				// get next stat
				cpuPrev = cpuCurrent
				netPrev = netCurrent
			}
		}
	}()
}

func (sa *SimpleAgent) Stop() {
	sa.conn.Close()
	sa.CancelFunc()
}

func StopSimpleAgent(uuid string) error {
	utils.Log.Debug(fmt.Sprintf("Stop collect data: %#v", SimpleAgents))

	agent, ok := SimpleAgents[uuid]
	if !ok {
		return errors.New("agent not found")
	}

	agent.Stop()
	delete(SimpleAgents, uuid)

	return nil
}

func (sa *SimpleAgent) Restart() error {
	conn, err := goph.NewConn(
		&goph.Config{
			User:     sa.User,
			Addr:     sa.Ip,
			Port:     sa.Port,
			Auth:     goph.Password(sa.Password),
			Callback: ssh.InsecureIgnoreHostKey(), // set unknown host key callback
		})
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not connect to vm: %#v", sa.AgentInfo), err)
		return err
	}

	sa.conn = conn

	return nil
}

func (sa *SimpleAgent) Status() {

}
