package agent

import (
	"context"
	"errors"
	"fmt"
	"kvm-dashboard/consts"
	"kvm-dashboard/model"
	"kvm-dashboard/utils"
	"kvm-dashboard/vm/data"
	"kvm-dashboard/vm/vm_utils"
	"time"

	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
)

var ProcessAgents map[string]*ProcessAgent

type ProcessAgent struct {
	*AgentInfo
	AgentInterface
	conn *goph.Client

	Ctx         context.Context // one ctx for one goroutine
	CancelFunc  context.CancelFunc
	ProcessData chan []*data.ProcessData
}

func NewProcessAgentWithJumpServer(agentInfo *AgentInfo, jumpServerInfo *model.Machine) (*ProcessAgent, error) {
	if ProcessAgents == nil {
		ProcessAgents = make(map[string]*ProcessAgent)
	}

	jumpServer, err := vm_utils.NewJumpServer(jumpServerInfo.Username,
		jumpServerInfo.Password, jumpServerInfo.Ip, jumpServerInfo.SshPort)
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not connect to jumpserver: %#v", jumpServerInfo), err)
		return nil, err
	}

	// use jump serve to get connection
	conn, err := vm_utils.JumpToServer(jumpServer,
		agentInfo.User, agentInfo.Password, agentInfo.Ip, fmt.Sprintf("%d", agentInfo.Port))
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not connect to vm: %#v", agentInfo), err)
		return nil, err
	}

	// create context
	ctx, cancelFunc := context.WithCancel(context.Background())

	agent := &ProcessAgent{
		AgentInfo:  agentInfo,
		conn:       conn,
		Ctx:        ctx,
		CancelFunc: cancelFunc,
	}

	return agent, nil
}

func NewProcessAgent(agentInfo *AgentInfo) (*ProcessAgent, error) {
	if ProcessAgents == nil {
		ProcessAgents = make(map[string]*ProcessAgent)
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

	agent := &ProcessAgent{
		AgentInfo:  agentInfo,
		conn:       conn,
		Ctx:        ctx,
		CancelFunc: cancelFunc,
	}

	return agent, nil
}

func (pa *ProcessAgent) Start(uuid string) {
	// store started agent
	ProcessAgents[uuid] = pa

	// defer pa.conn.Close()
	pa.ProcessData = make(chan []*data.ProcessData, 100) // set the buffer size to 100

	go func() {
		defer close(pa.ProcessData)

		for {
			select {
			case <-pa.Ctx.Done():
				utils.Log.Info(fmt.Sprintf("Stop process agent goroutine: %#v", uuid))
				return
			default:
				// delay for interval to get diff
				time.Sleep(time.Duration(consts.VM_DATA_REALTIME_INTERVAL) * time.Second)

				topInfo, err := vm_utils.GetCommandOutput(pa.conn, "top", "-b", "-n", "1") // batch mode, 1 iteration
				if err != nil {
					utils.Log.Error(fmt.Sprintf("Can not get top info on vm: %#v", pa.AgentInfo), err)
					return
				}
				processData := vm_utils.ParseTopData(topInfo)
				// utils.Log.Info(fmt.Sprintf("Process data: %#v", processData))
				pa.ProcessData <- processData

			}
		}
	}()
}

func (pa *ProcessAgent) Stop() {
	pa.conn.Close()
	pa.CancelFunc()
}

func StopProcessAgent(uuid string) error {
	utils.Log.Debug(fmt.Sprintf("Stop collect data: %#v", ProcessAgents))

	agent, ok := ProcessAgents[uuid]
	if !ok {
		return errors.New("agent not found")
	}

	agent.Stop()
	delete(ProcessAgents, uuid)

	return nil
}

func (pa *ProcessAgent) Restart() error {
	conn, err := goph.NewConn(
		&goph.Config{
			User:     pa.User,
			Addr:     pa.Ip,
			Port:     pa.Port,
			Auth:     goph.Password(pa.Password),
			Callback: ssh.InsecureIgnoreHostKey(), // set unknown host key callback
		})
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not connect to vm: %#v", pa.AgentInfo), err)
		return err
	}

	pa.conn = conn

	return nil
}

func (pa *ProcessAgent) Status() {

}
