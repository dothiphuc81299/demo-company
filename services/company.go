package services

import (
	"errors"
	
	"demo-company/dao"
	"demo-company/models"
	
)

// CompanyCreate ...
func CompanyCreate(payload models.CompanyCreatePayload) (models.CompanyBSON, error) {
	var (
		company = companyCreatePayloadToBSON(payload)
	)

	// company created
	doc, err := dao.CompanyCreate(company)
	

	// if err
	if err != nil {
		err = errors.New("Khong the tao company")
		return doc, err
	}

	return doc, err
}
