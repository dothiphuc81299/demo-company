package grpcuser

import (
	"sort"
	"sync"

	"demo-company/models"
	transactionpb "demo-company/proto/models/transaction"
	"demo-company/utils"
)

func convertToTransactionDetailList(data []*transactionpb.TransactionDetail) []models.TransactionDetail {
	var (
		result = make([]models.TransactionDetail, 0)
		wg     sync.WaitGroup
	)

	total := len(data)
	// Add process
	wg.Add(total)

	for index := range data {
		go func(index int) {
			defer wg.Done()

			// Convert to TransactionDetail
			transaction := convertToTransactionDetail(data[index])

			// Append
			result = append(result, transaction)
		}(index)
	}

	// Wait process
	wg.Wait()

	// Sort
	sort.Slice(result, func(i, j int) bool {
		return result[i].CreatedAt.After(result[j].CreatedAt)
	})

	return result
}

func convertToTransactionDetail(transaction *transactionpb.TransactionDetail) models.TransactionDetail {
	var (
		id       = utils.HelperParseStringToObjectID(transaction.Id)
		companyID = utils.HelperParseStringToObjectID(transaction.CompanyID)
		branchID  = utils.HelperParseStringToObjectID(transaction.BranchID)
		userID   = utils.HelperParseStringToObjectID(transaction.UserID)
		createdAt = utils.HelperConvertTimestampProtoToTime(transaction.CreatedAt)
	)

	// TransactionDetail
	result := models.TransactionDetail{
		ID:                     id,
		CompanyID:              companyID,
		BranchID:               branchID,
		UserID:                 userID,
		Amount:                 transaction.Amount,
		Commission:             transaction.Commission,
		CompanyCashbackPercent: transaction.CompanyCashbackPercent,
		CreatedAt:              createdAt,
	}
	return result
}
