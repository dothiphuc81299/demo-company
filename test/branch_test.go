package test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/xeipuuv/gojsonschema"
	"go.mongodb.org/mongo-driver/bson"

	"demo-company/apptest"
	"demo-company/models"
	"demo-company/modules/database"
	"demo-company/util"
)

type BranchSuite struct {
	suite.Suite
	e    *echo.Echo
	data models.BranchCreatePayload
}

func (suite *BranchSuite) SetupSuite() {
	// Init server ...
	suite.e = apptest.InitServer()

	// Clear data
	removeOldDataBranch()

	// Setup payload data
	suite.data = setupDataBranch()
}

func (suite *BranchSuite) TearDownSuite() {
	removeOldDataBranch()
}

func (suite *BranchSuite) TestBranchCreateSuccess() {
	var (
		payload      = suite.data
		response     util.Response
		schemaLoader = gojsonschema.NewReferenceLoader("file:///home/phuc/go/src/demo-company/schema/branch_create.json")
	)

	// Setup request
	req, _ := http.NewRequest(http.MethodPost, "/branches", util.HelperToIOReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	// Run HTTP server
	suite.e.ServeHTTP(res, req)

	// Parse
	json.Unmarshal([]byte(res.Body.String()), &response)

	// Create JSONLoader from go struct
	documentLoader := gojsonschema.NewGoLoader(response)

	// Validate json response
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

	// Test
	assert.Equal(suite.T(), true, result.Valid())
	assert.Equal(suite.T(), http.StatusOK, res.Code)
	assert.NotEqual(suite.T(), nil, response["data"])
}

func (suite *BranchSuite) TestBranchCreateFailureWithInvalidCompanyID() {
	var (
		payload = models.BranchCreatePayload{
			CompanyID: "08282882",
			Name:      suite.data.Name,
		}
		response util.Response
	)

	// Setup request
	req, _ := http.NewRequest(http.MethodPost, "/branches", util.HelperToIOReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	// Run HTTP server
	suite.e.ServeHTTP(res, req)

	// Parse ..
	json.Unmarshal([]byte(res.Body.String()), &response)

	// Test
	assert.Equal(suite.T(), http.StatusBadRequest, res.Code)
	assert.Equal(suite.T(), nil, response["data"])
}

func (suite *BranchSuite) TestBranchCreateFailureWithInvalidName() {
	var (
		payload = models.BranchCreatePayload{
			CompanyID: suite.data.CompanyID,
			Name:      "89",
		}
		response util.Response
	)

	// Setup request
	req, _ := http.NewRequest(http.MethodPost, "/branches", util.HelperToIOReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	// Run HTTP server
	suite.e.ServeHTTP(res, req)

	// Parse ..
	json.Unmarshal([]byte(res.Body.String()), &response)

	// Test
	assert.Equal(suite.T(), http.StatusBadRequest, res.Code)
	assert.Equal(suite.T(), nil, response["data"])
}

func (suite *BranchSuite) TestBranchCreateFailureWithNotFoundCompany() {
	var (
		payload = models.BranchCreatePayload{
			CompanyID: "5f24d45125ea51bc57a8285p",
			Name:      suite.data.Name,
		}
		response util.Response
	)

	// Setup request
	req, _ := http.NewRequest(http.MethodPost, "/branches", util.HelperToIOReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	// Run HTTP server
	suite.e.ServeHTTP(res, req)

	// Parse ..
	json.Unmarshal([]byte(res.Body.String()), &response)

	// Test
	assert.Equal(suite.T(), http.StatusBadRequest, res.Code)
	assert.Equal(suite.T(), nil, response["data"])
}

func TestBranchSuite(t *testing.T) {
	suite.Run(t, new(BranchSuite))
}

func setupDataBranch() models.BranchCreatePayload {
	var (
		companyIDString = util.HelperCompanyCreateFake()
	)
	payload := models.BranchCreatePayload{
		CompanyID: companyIDString,
		Name:      "89Nguyen chanh",
	}
	return payload
}

func removeOldDataBranch() {
	database.BranchCol().DeleteMany(context.Background(), bson.M{})
	database.CompanyCol().DeleteMany(context.Background(), bson.M{})
}
