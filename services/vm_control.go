package services

import (
	"kvm-dashboard/consts"
	"kvm-dashboard/utils"
	"kvm-dashboard/vm/agent"
)

func (s *Service) StartVM(uuid string) error {
	libvirtAgent, err := agent.NewLibvirtAgent(consts.LIBVIRT_URL)
	if err != nil {
		utils.Log.Error("Can not create libvirt agent", err)
		return err
	}

	err = libvirtAgent.StartVM(uuid)
	if err != nil {
		utils.Log.Error("Can not start vm", err)
		return err
	}

	return nil
}

func (s *Service) StopVM(uuid string) error {
	libvirtAgent, err := agent.NewLibvirtAgent(consts.LIBVIRT_URL)
	if err != nil {
		utils.Log.Error("Can not create libvirt agent", err)
		return err
	}

	err = libvirtAgent.StopVM(uuid)
	if err != nil {
		utils.Log.Error("Can not start vm", err)
		return err
	}

	return nil
}

func (s *Service) SuspendVM(uuid string) error {
	libvirtAgent, err := agent.NewLibvirtAgent(consts.LIBVIRT_URL)
	if err != nil {
		utils.Log.Error("Can not create libvirt agent", err)
		return err
	}

	err = libvirtAgent.SuspendVM(uuid)
	if err != nil {
		utils.Log.Error("Can not start vm", err)
		return err
	}

	return nil
}

func (s *Service) ResumeVM(uuid string) error {
	libvirtAgent, err := agent.NewLibvirtAgent(consts.LIBVIRT_URL)
	if err != nil {
		utils.Log.Error("Can not create libvirt agent", err)
		return err
	}

	err = libvirtAgent.ResumeVM(uuid)
	if err != nil {
		utils.Log.Error("Can not start vm", err)
		return err
	}

	return nil
}
