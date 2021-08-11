package service

import (
	"errors"
	"flag"
)

func Run() error {

	// Get service configuration from CLI inputs
	svcConf, err := parseConfig(NewServiceConfig())
	if err != nil {
		flag.Usage()
		return err
	}

	// Check access to SCM first, if it's allowed, get the handle
	// <code>

	// Create svc
	svc := Create(svcConf)
	if svc != nil {
		return err
	}
	return nil
}

// Parse Command Line flags
func parseConfig(svcConf *ServiceConfig) (*ServiceConfig, error) {
	flag.StringVar(&svcConf.Name, "name", "", "Service Name: TestSvc")
	flag.StringVar(&svcConf.ExePath, "exepath", "", "Executable Path: C:\\program.exe")
	// flag.StringVar(&svcConf.StartupType, "start", "demand", "Startup Type")
	// flag.StringVar(&svcConf.SvcType, "type", "own", "Service Type")
	flag.Parse()

	if svcConf.Name == "" || svcConf.ExePath == "" {
		return nil, errors.New("[Error] empty arguments")
	}
	return svcConf, nil
}

// func Create() {
// 	log := hclog.New(&hclog.LoggerOptions{
// 		Name:  "Install service",
// 		Level: hclog.LevelFromString("DEBUG"),
// 	})

// 	// Get a handle to the scm database
// 	scmHandle, err := NewSCMHandle(log)
// 	if err != nil {
// 		return
// 	}

// 	// Create a service
// 	svcHandle, err := scmHandle.CreateService()
// 	if err != nil {
// 		return
// 	}

// 	log.Info("Service has been created")

// 	// Start the service
// 	err = svcHandle.StartService()
// 	if err != nil {
// 		return
// 	}

// 	// Finally close the handlers
// 	err = svcHandle.Close()
// 	if err != nil {
// 		return
// 	}

// 	err = scmHandle.Close()
// 	if err != nil {
// 		return
// 	}
// }
