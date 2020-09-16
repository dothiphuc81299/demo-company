package grpcuser

import (
	"context"
	"log"
	"time"

	"demo-company/models"
	transactionpb "demo-company/proto/models/transaction"
)

// GetTransactionDetailByCompanyID ...
func GetTransactionDetailByCompanyID(companyID string) (transactions []models.TransactionDetail, err error) {
	// Setup client
	clientConn, client := CreateClient()
	defer clientConn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Call GetTransactionDetailByCompanyID
	result, err := client.GetTransactionDetailByCompanyID(ctx, &transactionpb.GetTransactionDetailByCompanyIDRequest{CompanyID: companyID})
	if err != nil {
		log.Printf("Call grpc get transaction by companyID error %v\n", err)
		return
	}
	
	// Convert to Company brief
	transactions = convertToTransactionDetailList(result.TransactionDetail)
	return
}
