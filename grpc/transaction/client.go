package grpcuser

import (
	"log"

	"google.golang.org/grpc"

	transactionpb "demo-company/proto/models/transaction"
	"demo-company/config"
)

// CreateClient ...
func CreateClient() (*grpc.ClientConn, transactionpb.TransactionServiceClient) {
	envVars := config.GetEnv()
	address := envVars.GRPCAddresses.Transaction + envVars.GRPCPorts.Transaction

	clientConn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("err while dial %v", err)
	}

	client := transactionpb.NewTransactionServiceClient(clientConn)

	return clientConn, client	
}
