package vm_utils

import (
	"fmt"
	"io"
	"kvm-dashboard/consts"
	"kvm-dashboard/utils"
	"os/exec"
	"testing"

	"github.com/magiconair/properties/assert"
	"libvirt.org/libvirt-go"
)

func TestExecCommand(t *testing.T) {
	cmd := exec.Command("ssh", "-t", "asklv@localhost")
	// cmd := exec.Command("ssh", "-t", "asklv@localhost", "sudo", "virsh", "console", "--force", "33ccd0be-a3a0-48ba-a5f7-b8d297cc7048")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		utils.Log.Error("Error:", err)
		return
	}
	// stream input
	stdin, err := cmd.StdinPipe()
	if err != nil {
		utils.Log.Error("Error:", err)
		return
	}

	if err = cmd.Start(); err != nil {
		utils.Log.Error(err)
		return
	}

	inputData := []string{consts.USERNAME, consts.PASSWORD, "pwd"}

	for _, data := range inputData {
		utils.Log.Info("222")

		_, err := io.WriteString(stdin, data+"\n")
		if err != nil {
			utils.Log.Error(err)
			return
		}
	}

	utils.Log.Info("111")
	output, err := io.ReadAll(stdout)
	if err != nil {
		utils.Log.Error(err)
	}

	fmt.Println("Console output:")
	fmt.Println(string(output))

	stdin.Close()

	if err := cmd.Wait(); err != nil {
		utils.Log.Error(err)
		return
	}
}

func TestAddArpEntry(t *testing.T) {
	base, err := NewJumpServer(
		"root",
		"",
		"127.0.0.1",
		"22",
	)
	assert.Equal(t, err, nil)

	err = AddArpEntry(base, "33ccd0be-a3a0-48ba-a5f7-b8d297cc7048")
	if err != nil {
		utils.Log.Error(err)
		t.Error(err)
	}
}

func TestVmLogin(t *testing.T) {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		utils.Log.Error(err)
		t.Error(err)
	}
	defer conn.Close()

	vmName := "vm1"
	vm, err := conn.LookupDomainByName(vmName)
	if err != nil {
		utils.Log.Error(err)
		t.Error(err)
	}
	defer vm.Free()

	// *Stream is in conn
	stream, err := conn.NewStream(libvirt.STREAM_NONBLOCK)
	if err != nil {
		utils.Log.Error(err)
		t.Error(err)
	}
	defer stream.Free()

	err = vm.OpenConsole("", stream, libvirt.DOMAIN_CONSOLE_FORCE)
	if err != nil {
		utils.Log.Error(err)
		t.Error(err)
	}

	stream.Send([]byte("root\n"))
	stream.Send([]byte("Ww1270278575\n"))
	// stream.Send([]byte("uname\n"))
	data := make([]byte, 1024)
	stream.Recv(data)
	utils.Log.Info(string(data))

	// utils.Log.Info("000")
	// // use block io send
	// err = stream.SendAll(func(s *libvirt.Stream, i int) ([]byte, error) {
	// 	return []byte("root\n"), nil
	// })
	// utils.Log.Info("111")
	// if err != nil {
	// 	utils.Log.Error(err)
	// 	t.Error(err)
	// }
	// utils.Log.Info("222")

	// // use block io recv
	// err = stream.RecvAll(func(s *libvirt.Stream, b []byte) (int, error) {
	// 	utils.Log.Info(string(b))
	// 	return -1, nil
	// })
	// utils.Log.Info("333")
	// if err != nil {
	// 	utils.Log.Error(err)
	// 	t.Error(err)
	// }
	// utils.Log.Info("444")

	// utils.Log.Info("000")
	// // add hook for stream readable event
	// stream.EventAddCallback(libvirt.STREAM_EVENT_WRITABLE, func(s *libvirt.Stream, events libvirt.StreamEventType) {
	// 	flag, err := stream.Send([]byte("root\n"))
	// 	if err != nil {
	// 		utils.Log.Error(err)
	// 		t.Error(err)
	// 	}
	// 	utils.Log.Info("111")

	// 	if flag < 0 {
	// 		utils.Log.Error("send error")
	// 		stream.Abort()
	// 	}
	// })
	// utils.Log.Info("222")

	// stream.EventAddCallback(libvirt.STREAM_EVENT_READABLE, func(s *libvirt.Stream, events libvirt.StreamEventType) {
	// 	buf := make([]byte, 1024)
	// 	data, err := s.Recv(buf)
	// 	if err != nil {
	// 		utils.Log.Error(err)
	// 	}
	// 	utils.Log.Info("333")
	// 	if data > 0 {
	// 		utils.Log.Info(string(buf[:data]))
	// 	}
	// })
	// utils.Log.Info("444")

	// for {
	// }

	// flag, err := stream.Send([]byte("root\n"))
	// if err != nil {
	// 	utils.Log.Error(err)
	// 	t.Error(err)
	// }
	// if flag < 0 {
	// 	utils.Log.Error("send error")
	// 	stream.Abort()
	// }

	// buf := make([]byte, 1024)
	// for i := 0; i < 10; i++ {
	// 	flag, err := stream.Send([]byte("root\n"))
	// 	if err != nil {
	// 		utils.Log.Error(err)
	// 		t.Error(err)
	// 	}
	// 	if flag < 0 {
	// 		utils.Log.Error("send error")
	// 		stream.Abort()
	// 	}

	// 	// time.Sleep(5 * time.Second)
	// 	stream.RecvAll(func(s *libvirt.Stream, b []byte) (int, error) {
	// 		utils.Log.Info(string(b))
	// 		return 0, nil
	// 	})
	// 	// n, err := stream.Recv(buf)
	// 	// utils.Log.Info(n)
	// 	// if err != nil {
	// 	// 	utils.Log.Error(err)
	// 	// }
	// 	// if n > 0 {
	// 	// 	utils.Log.Info(string(buf[:n]))
	// 	// }
	// }

	utils.Log.Info("end")
}
