package service

// ServiceConfig contains minimum configurations
// needed for creating a service
type ServiceConfig struct {
	Name    string
	ExePath string
	// SvcType     string
	// StartupType string
}

// Still idk if this one is really needed
// Since the config is generated from cli
// But I will just keep it first
func NewServiceConfig() ServiceConfig {
	return ServiceConfig{}
}
