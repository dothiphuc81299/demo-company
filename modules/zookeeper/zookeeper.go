package zookeeper

import (
	"fmt"
	"os"
	"time"

	"github.com/samuel/go-zookeeper/zk"

	"demo-company/config"
)

var conn *zk.Conn

// Connect ...
func Connect() {
	var (
		uri     = os.Getenv("ZOOKEEPER_URI")
		envVars = config.GetEnv()
	)
	conn, _, err := zk.Connect([]string{uri}, time.Second*30)
	
	if err != nil {
		fmt.Println("Can't connect to zookeeper", uri)
		panic(err)
	}
	fmt.Println("Connected to zookeeper", uri)

	// App Port
	appPort, _, _ := conn.Get("/app/port")
	envVars.AppPort = string(appPort)

	// Database
	databaseURI, _, _ := conn.Get("/database/uri")
	envVars.Database.URI = string(databaseURI)
	databaseName, _, _ := conn.Get("/database/name")
	envVars.Database.Name = string(databaseName)
	databaseTestName, _, _ := conn.Get("/database/test/name")
	envVars.Database.TestName = string(databaseTestName)

	// grpc Server
	grpcURI, _, _ := conn.Get("/grpc/uri")
	envVars.GRPC.URI = string(grpcURI)
}
