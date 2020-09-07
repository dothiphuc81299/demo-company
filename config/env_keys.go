package config
// ENV  ...
type  ENV struct {
	IsDev bool 

	ZookeeperURI string 
	ZookeeperTestURI string

	//Add port 
	CompanyPort string 

	Database struct {
		URI string 
		CompanyName string
		TestName string
	}

	GRPC struct {
		GRPCURI string
	}
}

var env ENV

// initENV ..
func initENV(){
	env =ENV{
		IsDev :true,
	}
}

// GetEnv ... 
func GetEnv() *ENV {
	return &env 
}