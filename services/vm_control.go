package services

import (
	"kvm-dashboard/utils"
	"kvm-dashboard/vm/agent"
)

func (svc *Service) StartVM(uuid string) error {
	libvirtURL := svc.GetMachineInfo(uuid).LibvirtUrl
	libvirtAgent, err := agent.NewLibvirtAgent(libvirtURL)
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

func (svc *Service) StopVM(uuid string) error {
	libvirtURL := svc.GetMachineInfo(uuid).LibvirtUrl
	libvirtAgent, err := agent.NewLibvirtAgent(libvirtURL)
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

func (svc *Service) SuspendVM(uuid string) error {
	libvirtURL := svc.GetMachineInfo(uuid).LibvirtUrl
	libvirtAgent, err := agent.NewLibvirtAgent(libvirtURL)
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

func (svc *Service) ResumeVM(uuid string) error {
	libvirtURL := svc.GetMachineInfo(uuid).LibvirtUrl
	libvirtAgent, err := agent.NewLibvirtAgent(libvirtURL)
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
