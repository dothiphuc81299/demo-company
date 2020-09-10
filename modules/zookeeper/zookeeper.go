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

	// App port
	appCompanyPort, _, _ := conn.Get("/app/port/company")
	envVars.AppPort = string(appCompanyPort)

	// Database
	databaseURI, _, _ := conn.Get("/database/uri")
	envVars.Database.URI = string(databaseURI)
	databaseCompanyName, _, _ := conn.Get("/database/name/company")
	envVars.Database.Name = string(databaseCompanyName)
	databaseTestName, _, _ := conn.Get("/database/test/name")
	envVars.Database.TestName = string(databaseTestName)

	// GRPCAddresses
	grpcAddressCompany, _, _ := conn.Get("/grpc/uri/company")
	envVars.GRPCAddresses.Company = string(grpcAddressCompany)
	grpcAddressTransaction, _, _ := conn.Get("/grpc/uri/transaction")
	envVars.GRPCAddresses.Transaction = string(grpcAddressTransaction)

	// GRPCPorts
	grpcPortCompany, _, _ := conn.Get("/grpc/port/company")
	envVars.GRPCPorts.Company = string(grpcPortCompany)
	grpcPortTransaction, _, _ := conn.Get("/grpc/port/transaction")
	envVars.GRPCPorts.Transaction = string(grpcPortTransaction)
}
