package vm_utils

import (
	"fmt"
	"kvm-dashboard/utils"
	"net"
	"strconv"

	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
)

// Use jump server (host machine) to get remote/local info
func NewJumpServer(user, password, ip, port string) (*goph.Client, error) {
	port_int, err := strconv.ParseInt(port, 10, 0)
	if err != nil {
		utils.Log.Error("Can not convert ssh port to int", err)
	}

	jumpConfig := &goph.Config{
		User:     user,
		Addr:     ip,
		Port:     uint(port_int),
		Auth:     goph.Password(password),
		Callback: ssh.InsecureIgnoreHostKey(), // set unknown host key callback
	}

	conn, err := goph.NewConn(jumpConfig)
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not connect to jumpserver: %#v", jumpConfig), err)
		return nil, err
	}

	return conn, nil
}

func JumpToServer(base *goph.Client, user, password, ip, port string) (*goph.Client, error) {
	port_int, err := strconv.ParseInt(port, 10, 0)
	if err != nil {
		utils.Log.Error("Can not convert ssh port to int", err)
	}

	newConfig := &goph.Config{
		User:     user,
		Addr:     ip,
		Port:     uint(port_int),
		Auth:     goph.Password(password),
		Callback: ssh.InsecureIgnoreHostKey(), // set unknown host key callback
	}

	newClient, err := jumpDial(base.Client, "tcp", newConfig)
	if err != nil {
		utils.Log.Error(fmt.Sprintf("Can not connect to dialserver: %#v", newConfig), err)
		return nil, err
	}

	// use new client
	base.Client = newClient
	base.Config = newConfig

	return base, nil
}

func jumpDial(base *ssh.Client, proto string, c *goph.Config) (*ssh.Client, error) {
	conn, err := base.Dial(proto, net.JoinHostPort(c.Addr, fmt.Sprint(c.Port)))
	if err != nil {
		return nil, err
	}

	ncc, chans, reqs, err := ssh.NewClientConn(conn, net.JoinHostPort(c.Addr, fmt.Sprint(c.Port)), &ssh.ClientConfig{
		User:            c.User,
		Auth:            c.Auth,
		Timeout:         c.Timeout,
		HostKeyCallback: c.Callback,
		BannerCallback:  c.BannerCallback,
	})
	if err != nil {
		return nil, err
	}

	sClient := ssh.NewClient(ncc, chans, reqs)
	return sClient, nil
}
