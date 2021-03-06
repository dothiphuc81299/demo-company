package services

import (
	"errors"

	"demo-company/dao"
	grpctransaction "demo-company/grpc/transaction"
	"demo-company/models"
)

// CompanyCreate ...
func CompanyCreate(payload models.CompanyCreatePayload) (models.CompanyBSON, error) {
	var (
		company = payload.ConvertToBSON()
	)

	// Create company
	doc, err := dao.CompanyCreate(company)

	// If err
	if err != nil {
		err = errors.New("khong the tao company")
		return doc, err
	}
	return doc, err
}

// TransactionFindByCompanyID ...
func TransactionFindByCompanyID(companyID string) ([]models.TransactionDetail, error) {
	var (
		result = make([]models.TransactionDetail, 0)
	)

	// Call gRPC get Transactions
	result, err := grpctransaction.GetTransactionDetailByCompanyID(companyID)
	if err != nil {
		err = errors.New(err.Error())
		return result, err
	}
	return result, err
}
