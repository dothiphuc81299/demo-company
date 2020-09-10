package grpcnode

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson"

	"demo-company/dao"
	companypb "demo-company/proto/models/company"
	"demo-company/util"
)

func getCompanyBriefByID(companyIDString string) (*companypb.GetCompanyBriefByIDResponse, error) {
	var (
		companyID = util.HelperParseStringToObjectID(companyIDString)
	)

	// Find Company
	company, err := dao.CompanyFindByID(companyID)
	if err != nil {
		err = errors.New("Not Found Company by ID")
		return nil, err
	}

	// Success
	result := &companypb.GetCompanyBriefByIDResponse{
		CompanyBrief: &companypb.CompanyBrief{
			Id:               companyIDString,
			Name:             company.Name,
			CashbackPercent:  company.CashbackPercent,
			TotalTransaction: company.TotalTransaction,
			TotalRevenue:     company.TotalRevenue,
		},
	}
	return result, nil
}

func getBranchBriefByID(branchIDString string) (*companypb.GetBranchBriefByIDResponse, error) {
	var (
		branchID = util.HelperParseStringToObjectID(branchIDString)
	)

	// Find Branch
	branch, err := dao.BranchFindByID(branchID)
	if err != nil {
		err = errors.New("Not Found Branch by ID")
		return nil, err
	}

	// Success
	result := &companypb.GetBranchBriefByIDResponse{
		BranchBrief: &companypb.BranchBrief{
			Id:               branchIDString,
			Name:             branch.Name,
			TotalTransaction: branch.TotalTransaction,
			TotalRevenue:     branch.TotalRevenue,
		},
	}
	return result, nil
}

func updateCompanyStatsByID(companyIDString string, totalTransaction int64, totalRevenue float64) (*companypb.UpdateCompanyStatsByIDResponse, error) {
	var (
		companyID = util.HelperParseStringToObjectID(companyIDString)
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
		err = errors.New("Update CompanyStats error")
		return nil, err
	}

	// Success
	result := &companypb.UpdateCompanyStatsByIDResponse{}
	return result, nil
}

func updateBranchStatsByID(branchIDString string, totalTransaction int64, totalRevenue float64) (*companypb.UpdateBranchStatsByIDResponse, error) {
	var (
		branchID = util.HelperParseStringToObjectID(branchIDString)
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
		err = errors.New("Update BranchStats error")
		return nil, err
	}

	// Success
	result := &companypb.UpdateBranchStatsByIDResponse{}
	return result, nil
}
