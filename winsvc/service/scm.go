package service

import (
	"fmt"

	"github.com/fahmifj/winapi-go/winsvc/util"
	"golang.org/x/sys/windows"
)

// Access Rights for the Service Control Manager
// (see: https://docs.microsoft.com/en-us/windows/win32/services/service-security-and-access-rights?redirectedfrom=MSDN#access_rights_for_the_service_control_manager)

type ISCM interface {
	// Connect() error
	Create(*ServiceConfig) error
}

// SCM is a Service Control Manager(SCM)
// (see: https://docs.microsoft.com/en-us/windows/win32/api/winsvc/nf-winsvc-openscmanagerw)
// type SCM struct {
// 	Handle windows.Handle
// }

// Create is create 🔨
func CreateStart(conf *ServiceConfig) error {

	// Note: this should be separated
	// Connect to SCM database and get a handle with desired access
	scmHandle, err := windows.OpenSCManager(nil, nil, windows.SC_MANAGER_ALL_ACCESS)
	if err != nil {
		fmt.Println("[Error] Failed to connect to the SCM database: ", err)
		return err
	}

	// Check if the service exists
	// <code>

	// Create service
	svcHandle, err := windows.CreateService(
		scmHandle,
		util.StringToUTF16Ptr(conf.Name),
		nil,
		windows.SERVICE_ALL_ACCESS,        // https://docs.microsoft.com/en-us/windows/win32/services/service-security-and-access-rights?redirectedfrom=MSDN#access-rights-for-a-service
		windows.SERVICE_WIN32_OWN_PROCESS, //
		windows.SERVICE_DEMAND_START,      //
		windows.SERVICE_ERROR_NORMAL,      //
		util.StringToUTF16Ptr(conf.ExePath),
		nil,
		nil,
		nil,
		nil,
		nil,
	)
	if err != nil {
		fmt.Printf("[Error] Failed to create service: %v", err)
		return err
	}

	err = windows.StartService(svcHandle, 0, nil)
	if err != nil {
		return err
	}

	if err := windows.CloseServiceHandle(svcHandle); err != nil {
		// failed to close the handle
		return err
	}
	if err := windows.CloseHandle(scmHandle); err != nil {
		// failed to close the handle
		return err
	}

	fmt.Printf("[Info] Service %v has been installed", conf.Name)
	// If all ok, return nil
	return nil
}
