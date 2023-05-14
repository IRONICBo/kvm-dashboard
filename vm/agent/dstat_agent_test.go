package agent

import (
	"bufio"
	"kvm-dashboard/utils"
	"os/exec"
	"testing"
	"time"
)

func TestDstatCommand(t *testing.T) {
	// cmd := exec.Command("dstat", "-l -c --proc --proc-count -y -i --ipc -m --vm -g -s -n --net-packets --socket --raw --tcp --udp --unix -d -r --aio --disk-tps --disk-util --fs --lock")
	cmd := exec.Command("dstat", "-l", "-c", "--proc", "--proc-count", "-y", "-i", "--ipc", "-m", "--vm", "-g", "-s", "-n", "--net-packets", "--socket", "--raw", "--tcp", "--udp", "--unix", "-d", "-r", "--aio", "--disk-tps", "--disk-util", "--fs", "--lock")
	out, err := cmd.StdoutPipe() // start a pipe to get stdout

	if err != nil {
		return
	}

	go func() {
		cmd.Start()
		scanner := bufio.NewScanner(out)

		for scanner.Scan() {
			t.Log(scanner.Text())
		}
	}()

	time.Sleep(2 * time.Second)
}

func TestDstatAgent(t *testing.T) {
	// dstatAgent := &DstatAgent{
	// 	AgentInfo: &AgentInfo{
	// 		UUID:     "test",
	// 		User:     "root",
	// 		Password: "Ia9nf5JUfWyVhtHQ",
	// 		Ip:       "127.0.0.1",
	// 		Port:     22,
	// 	},
	// }

	dstatAgent := &DstatAgent{
		AgentInfo: &AgentInfo{
			UUID:     "test",
			User:     "root",
			Password: "123123",
			Ip:       "192.168.122.166",
			Port:     22,
		},
	}

	utils.LogWithInfo("Start dstat agent")

	dstatAgent.Start()
	// time.Sleep(10 * time.Second)
}
