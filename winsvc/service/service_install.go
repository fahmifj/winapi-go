package service

import "github.com/hashicorp/go-hclog"

// TODO: 🙃

func SvcInstall() {
	log := hclog.New(&hclog.LoggerOptions{
		Name:  "Install service",
		Level: hclog.LevelFromString("DEBUG"),
	})

	// Get a handle to the scm database
	scmHandle, err := NewSCMHandle(log)
	if err != nil {
		panic(err)
	}

	// Create a service
	svcHandle, err := scmHandle.CreateService()
	if err != nil {
		panic(err)
	}

	// Start the service
	err = svcHandle.StartService()
	if err != nil {
		panic(err)
	}

	// Finally close the handlers
	err = svcHandle.Close()
	if err != nil {
		panic(err)
	}

	err = scmHandle.Close()
	if err != nil {
		panic(err)
	}
}
