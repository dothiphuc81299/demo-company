package services

import (
	"errors"

	"demo-company/dao"
	"demo-company/models"
)

// BranchCreate ...
func BranchCreate(payload models.BranchCreatePayload) (models.BranchBSON, error) {
	var (
		branch = branchCreatePayloadToBSON(payload)
	)

	// Create branch
	doc, err := dao.BranchCreate(branch)

	// if err
	if err != nil {
		err = errors.New("Khong the tao branch")
		return doc, err
	}
	return doc, err
}
