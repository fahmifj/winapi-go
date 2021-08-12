package service

import (
	"errors"
	"flag"
)

func Run() error {

	// Get service configuration from CLI inputs
	conf := NewServiceConfig()
	svc, err := parseConfig(&conf)
	if err != nil {
		flag.Usage()
		return err
	}

	// Check access to SCM first, if it's allowed, get the handle
	// <code>

	// Create svc
	err = CreateStart(svc)
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
