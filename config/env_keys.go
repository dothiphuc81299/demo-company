package config

// ENV  ...
type ENV struct {
	IsDev bool

	// ZookeeperURI
	ZookeeperURI string

	//Add port
	AppPort string

	Database struct {
		URI      string
		Name     string
		TestName string
	}

	// gRPC addresses
	GRPCAddresses struct {
		Company     string
		Transaction string
	}

	// gRPC ports
	GRPCPorts struct {
		Company     string
		Transaction string
	}
}

var env ENV

// InitENV ..
func InitENV() {
	env = ENV{
		IsDev: true,
	}
}

// GetEnv ...
func GetEnv() *ENV {
	return &env
}
