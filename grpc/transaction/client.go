package grpcuser

import (
	"log"

	"google.golang.org/grpc"

	"demo-company/config"
	transactionpb "demo-company/proto/models/transaction"
)

// CreateClient ...
func CreateClient() (*grpc.ClientConn, transactionpb.TransactionServiceClient) {
	envVars := config.GetEnv()
	address := envVars.GRPCAddresses.Transaction + envVars.GRPCPorts.Transaction

	// Create a client connection
	clientConn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("err while dial %v", err)
	}

	// Create transaction service
	client := transactionpb.NewTransactionServiceClient(clientConn)
	return clientConn, client
}
