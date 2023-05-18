package agent

import (
	"bufio"
	"fmt"
	"kvm-dashboard/utils"
	"kvm-dashboard/vm/data"

	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
)

type DstatAgent struct {
	*AgentInfo     // agent info
	AgentInterface // agent interface

	DstatData chan *data.DstatData // use go routine to collect dstat data
}

func NewDstatAgent(agentInfo *AgentInfo) *DstatAgent {
	return &DstatAgent{
		AgentInfo: agentInfo,
		DstatData: make(chan *data.DstatData, 100), // set the buffer size to 100
	}
}

func (da *DstatAgent) Start() {
	// connect to vm
	// client, err := goph.New(da.AgentInfo.User, da.AgentInfo.Ip, goph.Password(da.AgentInfo.Password))
	client, err := goph.NewConn(
		&goph.Config{
			User:     da.User,
			Addr:     da.Ip,
			Port:     da.Port,
			Auth:     goph.Password(da.Password),
			Callback: ssh.InsecureIgnoreHostKey(), // set unknown host key callback
		})
	// defer client.Close()

	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not connect to vm: %#v", da.AgentInfo), err)
		return
	}

	// run dstat
	// cmd, err := client.Command("ls", "-a")
	cmd, err := client.Command("dstat", data.DstatParam...)
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not run dstat on vm: %#v", da.AgentInfo), err)
		return
	}

	// commmand runs forever, so need pipe out
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not get output of dstat on vm: %#v", da.AgentInfo), err)
		return
	}
	err = cmd.Start()
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not run dstat on vm: %#v", da.AgentInfo), err)
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

func (da *DstatAgent) Stop() {
}

func (da *DstatAgent) Restart() {
}

func (da *DstatAgent) Status() {
}
