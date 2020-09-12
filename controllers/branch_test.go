package controllers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"

	"demo-company/util"
	"demo-company/models"
	"demo-company/modules/database"
)

type BranchSuite struct {
	suite.Suite
	e *echo.Echo
}

func (suite BranchSuite) SetupSuite() {
	suite.e =InitEcho()
	util.HelperCompanyCreateFake()
	RemoveOldDataCompany()
	removeOldDataBranch()
}

func (suite BranchSuite) TearDownSuite() {
	RemoveOldDataCompany()
	removeOldDataBranch()
}

func removeOldDataBranch() {
	database.BranchCol().DeleteMany(context.Background(),bson.M{})
}

func (suite *BranchSuite) TestSuccess(){
	var (
		payload =  models.BranchCreatePayload {
			CompanyID : util.CompanyStringSuccess,
			Name: "89Nguyen chanh",
		}
		response util.Response
	)
	suite.e = InitEcho()

	//set up request
	req, _ := http.NewRequest(http.MethodPost, "/branches", util.HelperToIOReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	
	suite.e.ServeHTTP(res, req)

	//parse .. 
	json.Unmarshal([]byte(res.Body.String()), &response)	
	
	//Test
	assert.Equal(suite.T(), http.StatusOK, res.Code)
	assert.NotEqual(suite.T(), nil, response["data"])
	assert.Equal(suite.T(), "thanh cong!", response["message"])
}

func (suite *BranchSuite) TestInvalidCompanyIDFail(){
	var (
		payload =  models.BranchCreatePayload {
			CompanyID : util.CompanyStringInvalid,
			Name: "89Nguyen chanh",
		}
		response util.Response
	)
	suite.e = InitEcho()

	//set up request
	req, _ := http.NewRequest(http.MethodPost, "/branches", util.HelperToIOReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	
	suite.e.ServeHTTP(res, req)

	//parse .. 
	json.Unmarshal([]byte(res.Body.String()), &response)	
	
	//Test
	assert.Equal(suite.T(), http.StatusBadRequest, res.Code)
	assert.Equal(suite.T(), nil, response["data"])
	assert.NotEqual(suite.T(), "thanh cong!", response["message"])
}

func (suite *BranchSuite) TestInvalidNameFail(){
	var (
		payload =  models.BranchCreatePayload {
			CompanyID : util.CompanyStringSuccess,
			Name: "89",
		}
		response util.Response
	)
	suite.e = InitEcho()

	//set up request
	req, _ := http.NewRequest(http.MethodPost, "/branches", util.HelperToIOReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	
	suite.e.ServeHTTP(res, req)

	//parse .. 
	json.Unmarshal([]byte(res.Body.String()), &response)	
	
	//Test
	assert.Equal(suite.T(), http.StatusBadRequest, res.Code)
	assert.Equal(suite.T(), nil, response["data"])
	assert.NotEqual(suite.T(), "thanh cong!", response["message"])
}

func TestBranchSuite(t *testing.T) {
	suite.Run(t, new(BranchSuite))
}

