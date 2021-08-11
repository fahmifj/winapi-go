package service

import (
	"fmt"

	"golang.org/x/sys/windows"
)

type IService interface {
	Start() error
}

// Service represents a Service
// It can control a service depending on
// the access given during its creation.
type Service struct {
	Handle windows.Handle
}

func (s *Service) Start() error {

	// connect to scm to get a handle
	// <code>

	// use the handle to open the service
	// <code>

	// query the service if it's already running
	// <code>

	// start the service
	// s.Log.Debug("Attempting to start the service using", s.Handle)
	err := windows.StartService(s.Handle, 0, nil)
	if err != nil {
		fmt.Print("Failed to start the service", err)
		return err
	}

	windows.CloseServiceHandle(s.Handle)
	return nil
}
