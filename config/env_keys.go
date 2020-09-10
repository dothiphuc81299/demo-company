package config
// ENV  ...
type  ENV struct {
	IsDev bool 

	ZookeeperURI string 
	ZookeeperTestURI string

	//Add port 
	AppPort string 

	Database struct {
		URI string 
		Name string
		TestName string
	}

	GRPC struct {
		URI string
	}
}

var env ENV

// InitENV ..
func InitENV(){
	env =ENV{
		IsDev :true,
	}
}

// GetEnv ... 
func GetEnv() *ENV {
	return &env 
}