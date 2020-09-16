package grpcnode

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"demo-company/config"
	companypb "demo-company/proto/models/company"
)

// Node ...
type Node struct{}

// GetCompanyBriefByID ...
func (s *Node) GetCompanyBriefByID(ctx context.Context, req *companypb.GetCompanyBriefByIDRequest) (*companypb.GetCompanyBriefByIDResponse, error) {
	var (
		companyID = req.GetCompanyID()
	)

	// Get Company by id
	data, err := getCompanyBriefByID(companyID)

	result := &companypb.GetCompanyBriefByIDResponse{
		CompanyBrief: data,
	}
	return result, err
}

// GetBranchBriefByID ...
func (s *Node) GetBranchBriefByID(ctx context.Context, req *companypb.GetBranchBriefByIDRequest) (*companypb.GetBranchBriefByIDResponse, error) {
	var (
		branchID = req.GetBranchID()
	)

	// Get Branch by id
	data, err := getBranchBriefByID(branchID)

	result := &companypb.GetBranchBriefByIDResponse{
		BranchBrief: data,
	}
	return result, err
}

// UpdateCompanyStatsByID ...
func (s *Node) UpdateCompanyStatsByID(ctx context.Context, req *companypb.UpdateCompanyStatsByIDRequest) (*companypb.UpdateCompanyStatsByIDResponse, error) {
	var (
		companyID        = req.GetId()
		totalTransaction = req.GetTotalTransaction()
		totalRevenue     = req.GetTotalRevenue()
	)

	// Update CompanyStats
	err := updateCompanyStatsByID(companyID, totalTransaction, totalRevenue)

	result := &companypb.UpdateCompanyStatsByIDResponse{}
	return result, err
}

// UpdateBranchStatsByID ...
func (s *Node) UpdateBranchStatsByID(ctx context.Context, req *companypb.UpdateBranchStatsByIDRequest) (*companypb.UpdateBranchStatsByIDResponse, error) {
	var (
		branchID         = req.GetId()
		totalTransaction = req.GetTotalTransaction()
		totalRevenue     = req.GetTotalRevenue()
	)

	// Update BranchStats
	err := updateBranchStatsByID(branchID, totalTransaction, totalRevenue)

	result := &companypb.UpdateBranchStatsByIDResponse{}
	return result, err
}

// Start ...
func Start() {
	envVars := config.GetEnv()
	companyPort := envVars.GRPCPorts.Company

	// Create Listen
	lis, err := net.Listen("tcp", companyPort)
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	// Create Service Server
	s := grpc.NewServer()
	companypb.RegisterCompanyServiceServer(s, &Node{})

	// Start Server
	log.Println(" gRPC server started on port:" + companyPort)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("err while %v", err)
	}
}
