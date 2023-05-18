package agent

import (
	"bufio"
	"fmt"
	"kvm-dashboard/utils"
	"kvm-dashboard/vm/data"

	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
)

// perf stat -I 1000 -x "|" -e branch-misses
type PerfAgent struct {
	*AgentInfo     // agent info
	AgentInterface // agent interface

	PerfData chan *data.PerfData // use go routine to collect dstat data
}

func NewPerfAgent(agentInfo *AgentInfo) *PerfAgent {
	return &PerfAgent{
		AgentInfo: agentInfo,
		PerfData:  make(chan *data.PerfData, 100), // set the buffer size to 100
	}
}

func (pa *PerfAgent) Start() {
	// connect to vm
	// todo: Port is not used in goph
	// client, err := goph.New(da.AgentInfo.User, da.AgentInfo.Ip, goph.Password(da.AgentInfo.Password))
	client, err := goph.NewConn(
		&goph.Config{
			User:     pa.User,
			Addr:     pa.Ip,
			Port:     pa.Port,
			Auth:     goph.Password(pa.Password),
			Callback: ssh.InsecureIgnoreHostKey(), // set unknown host key callback
		})
	// defer client.Close()

	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not connect to vm: %#v", pa.AgentInfo), err)
		return
	}

	// run dstat
	// cmd, err := client.Command("ls", "-a")
	// cmd, err := client.Command("perf", "stat", "-I 5000", "-x \"|\"", "-e", "branch-misses")
	cmd, err := client.Command("perf", "stat", "-I 1000", "-x \"|\"", "-e", data.PerfEvents)
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not run perf on vm: %#v", pa.AgentInfo), err)
		return
	}

	// utils.Log.Info(cmd.)

	// commmand runs forever, so need pipe out
	stdout, err := cmd.StderrPipe()
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not get output of perf on vm: %#v", pa.AgentInfo), err)
		return
	}
	err = cmd.Start()
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not run perf on vm: %#v", pa.AgentInfo), err)
		return
	}

	// parse dstat data
	go func() {
		scanner := bufio.NewScanner(stdout)

		for scanner.Scan() {
			line := scanner.Text()
			utils.Log.Info(line)
			// dstatData := utils.ParseDstatData(line)
			// da.DstatData <- dstatData
		}
	}()
}

func (pa *PerfAgent) Stop() {
}

func (pa *PerfAgent) Restart() {
}

func (pa *PerfAgent) Status() {
}
