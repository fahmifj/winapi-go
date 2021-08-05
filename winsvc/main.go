package main

import "github.com/fahmifj/winapi-go/winsvc/service"

func main() {
	// reference https://docs.microsoft.com/en-us/windows/win32/services/installing-a-service
	service.SvcInstall()
}
