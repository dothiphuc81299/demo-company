package grpcnode

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson"

	"demo-company/dao"
	companypb "demo-company/proto/models/company"
	"demo-company/util"
)

func getCompanyBriefByID(companyIDString string) (*companypb.CompanyBrief, error) {
	var (
		companyID,_ = util.HelperParseStringToObjectID(companyIDString)
	)

	// Find Company
	company, err := dao.CompanyFindByID(companyID)
	if err != nil {
		err = errors.New("not Found Company by ID")
		return nil, err
	}

	// Success
	result := &companypb.CompanyBrief{
		Id:               companyIDString,
		Name:             company.Name,
		CashbackPercent:  company.CashbackPercent,
		TotalTransaction: company.TotalTransaction,
		TotalRevenue:     company.TotalRevenue,
	}
	return result, nil
}

func getBranchBriefByID(branchIDString string) (*companypb.BranchBrief, error) {
	var (
		branchID,_ = util.HelperParseStringToObjectID(branchIDString)
	)

	// Find Branch
	branch, err := dao.BranchFindByID(branchID)
	if err != nil {
		err = errors.New("not Found Branch by ID")
		return nil, err
	}

	// Success
	result := &companypb.BranchBrief{
		Id:               branchIDString,
		Name:             branch.Name,
		TotalTransaction: branch.TotalTransaction,
		TotalRevenue:     branch.TotalRevenue,
	}
	return result, nil
}

func updateCompanyStatsByID(companyIDString string, totalTransaction int64, totalRevenue float64) error {
	var (
		companyID,_ = util.HelperParseStringToObjectID(companyIDString)
	)

	// Set filter and update
	filter := bson.M{"_id": companyID}
	update := bson.M{"$set": bson.M{
		"totalTransaction": totalTransaction,
		"totalRevenue":     totalRevenue,
	}}

	// Update Company
	err := dao.CompanyUpdateByID(filter, update)
	if err != nil {
		err = errors.New("update CompanyStats error")
		return err
	}
	return nil
}

func updateBranchStatsByID(branchIDString string, totalTransaction int64, totalRevenue float64) error {
	var (
		branchID,_ = util.HelperParseStringToObjectID(branchIDString)
	)

	// Set filter and update
	filter := bson.M{"_id": branchID}
	update := bson.M{"$set": bson.M{
		"totalTransaction": totalTransaction,
		"totalRevenue":     totalRevenue,
	}}

	// Update Branch
	err := dao.BranchUpdateByID(filter, update)
	if err != nil {
		err = errors.New("update BranchStats error")
		return err
	}
	return nil
}
