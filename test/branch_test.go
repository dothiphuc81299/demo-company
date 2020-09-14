package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/xeipuuv/gojsonschema"

	"demo-company/util"
	"demo-company/models"
	"demo-company/modules/database"
	"demo-company/apptest"
)

type BranchSuite struct {
	suite.Suite
	e *echo.Echo
	data models.BranchCreatePayload 
}

func (suite *BranchSuite) SetupSuite() {
	// init server ... 
	suite.e =apptest.InitServer()

	// clear data 
	RemoveOldDataCompany()
	removeOldDataBranch()
	
	// set up payload data 
	companyIDString := util.HelperCompanyCreateFake()
	suite.data = models.BranchCreatePayload{
			CompanyID: companyIDString,
			Name: "89Nguyen chanh",
	}
}

func (suite *BranchSuite) TearDownSuite() {
	RemoveOldDataCompany()
	removeOldDataBranch()
}

func removeOldDataBranch() {
	database.BranchCol().DeleteMany(context.Background(),bson.M{})
}

func (suite *BranchSuite) TestBranchCreateSuccess(){
	var (
		payload = suite.data
		response util.Response
		schemaLoader = gojsonschema.NewReferenceLoader("file:///home/phuc/go/src/demo-company/schema/branch_create.json")
	)
	
	//set up request
	req, _ := http.NewRequest(http.MethodPost, "/branches", util.HelperToIOReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	
	// Run HTTP server
	suite.e.ServeHTTP(res, req)

	//parse .. 
	json.Unmarshal([]byte(res.Body.String()), &response)

	documentLoader := gojsonschema.NewGoLoader(response)
	// validate ..
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
	//Test
	assert.Equal(suite.T(), http.StatusOK, res.Code)
	assert.NotEqual(suite.T(), nil, response["data"])
	assert.Equal(suite.T(), "thanh cong!", response["message"])
}

func (suite *BranchSuite) TestBranchCreateFailureWithInvalidCompanyID(){
	var (
		payload =  models.BranchCreatePayload {
			CompanyID : "08282882",
			Name: "89Nguyen chanh",
		}
		response util.Response
	)
	
	//set up request
	req, _ := http.NewRequest(http.MethodPost, "/branches", util.HelperToIOReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	
	// Run HTTP server
	suite.e.ServeHTTP(res, req)

	//parse .. 
	json.Unmarshal([]byte(res.Body.String()), &response)	
	
	//Test
	assert.Equal(suite.T(), http.StatusBadRequest, res.Code)
	assert.Equal(suite.T(), nil, response["data"])
	assert.NotEqual(suite.T(), "thanh cong!", response["message"])
}

func (suite *BranchSuite) TestBranchCreateWithFailureInvalidName(){
	var (
		data = suite.data 
		payload =  models.BranchCreatePayload {
			CompanyID : data.CompanyID, 
			Name: "89",
		}
		response util.Response
	)
	
	//set up request
	req, _ := http.NewRequest(http.MethodPost, "/branches", util.HelperToIOReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	
	// Run HTTP server 
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

