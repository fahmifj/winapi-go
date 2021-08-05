package main

// func main() {
// 	// log := hclog.New(&hclog.LoggerOptions{
// 	// 	Name:  "win-api",
// 	// 	Level: hclog.LevelFromString("DEBUG"),
// 	// })

// 	// connect to scm to get a handle to create a service
// 	scmHandle, err := windows.OpenSCManager(nil, nil, windows.SC_MANAGER_CREATE_SERVICE)
// 	if err != nil {
// 		fmt.Print("Error cannot create a service", err)
// 		os.Exit(-1)
// 	}
// 	// define service

// 	serviceName, _ := windows.UTF16PtrFromString("NcatService")
// 	serviceExecPath, _ := windows.UTF16PtrFromString(`C:\test-service\ncat.exe localhost 9000`)
// 	serviceRunOnBehalf, _ := windows.UTF16PtrFromString(`NT Authority\System`)

// 	// create a service
// 	svcHandle, err := windows.CreateService(
// 		scmHandle,
// 		serviceName,
// 		nil,
// 		windows.SERVICE_ALL_ACCESS,
// 		windows.SERVICE_WIN32_OWN_PROCESS,
// 		windows.SERVICE_DEMAND_START,
// 		windows.SERVICE_ERROR_NORMAL,
// 		serviceExecPath,
// 		nil,
// 		nil,
// 		nil,
// 		serviceRunOnBehalf,
// 		nil,
// 	)
// 	if err != nil {
// 		fmt.Print("err", err)
// 		os.Exit(-1)
// 	}

// 	// start the service
// 	if windows.StartService(svcHandle, 0, nil); err != nil {
// 		fmt.Print("Error Starting service", err)
// 		os.Exit(-1)
// 	}

// }
