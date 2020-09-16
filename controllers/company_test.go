package controllers_test

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
	"demo-company/utils"
)

type CompanyCreateSuite struct {
	suite.Suite
	e    *echo.Echo
	data models.CompanyCreatePayload
}

func (suite *CompanyCreateSuite) SetupSuite() {
	// Init server ...
	suite.e = apptest.InitServer()

	// Clear data
	removeOldDataCompany()

	// Setup payload data
	suite.data = models.CompanyCreatePayload{
		Name:            "UserTest",
		CashbackPercent: 19.2,
	}
}

func (suite *CompanyCreateSuite) TearDownSuite() {
	removeOldDataCompany()
}

func (suite *CompanyCreateSuite) TestCompanyCreateSuccess() {
	var (
		response     utils.Response
		payload      = suite.data
		schemaLoader = gojsonschema.NewReferenceLoader("file:///home/phuc/go/src/demo-company/schemas/company_create.json")
	)
	
	// Setup request
	req, err := http.NewRequest(http.MethodPost, "/companies", utils.HelperToIOReader(payload))
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

func (suite *CompanyCreateSuite) TestCompanyCreateFailureWithInvalidName() {
	var (
		response utils.Response
		payload  = models.CompanyCreatePayload{
			Name:            "",
			CashbackPercent: suite.data.CashbackPercent,
		}
	)

	// Setup request
	req, _ := http.NewRequest(http.MethodPost, "/companies", utils.HelperToIOReader(payload))
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

func (suite *CompanyCreateSuite) TestCompanyCreateFailureWithCashbackPercent() {
	var (
		response utils.Response
		payload  = models.CompanyCreatePayload{
			Name:            suite.data.Name,
			CashbackPercent: 0,
		}
	)

	// Setup request
	req, _ := http.NewRequest(http.MethodPost, "/companies", utils.HelperToIOReader(payload))
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

type TransactionFindByCompanyIDSuite struct {
	suite.Suite
	e        *echo.Echo
	paramURL string
}

func (suite *TransactionFindByCompanyIDSuite) SetupSuite() {
	// Setup server
	suite.e = apptest.InitServer()

	// Clear data
	removeOldDataCompany()

	// Setup data
	suite.paramURL = utils.HelperCompanyCreateFake()
}

func (suite *TransactionFindByCompanyIDSuite) TearDownSuite() {
	removeOldDataCompany()
}

func (suite *TransactionFindByCompanyIDSuite) TestTransactionFindByCompanyIDSuccess() {
	var (
		response     utils.Response
		schemaLoader = gojsonschema.NewReferenceLoader("file:///home/phuc/go/src/demo-company/schemas/company_create.json")
	)

	// Setup request
	url := "/companies/" + suite.paramURL + "/transactions"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// Run HTTP server
	suite.e.ServeHTTP(rec, req)

	// Parse
	json.Unmarshal([]byte(rec.Body.String()), &response)

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
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
}

func (suite *TransactionFindByCompanyIDSuite) TestTransactionFindByCompanyIDFailureWithInvalidCompanyID() {
	var (
		response utils.Response
		paramURL = "123"
	)

	// Setup request
	url := "/companies/" + paramURL + "/transactions"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// Run HTTP server
	suite.e.ServeHTTP(rec, req)

	// Parse
	json.Unmarshal([]byte(rec.Body.String()), &response)

	// Test
	assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
	assert.Equal(suite.T(), nil, response["data"])
}

func (suite *TransactionFindByCompanyIDSuite) TestTransactionFindByCompanyIDFailureWithNotFoundCompanyID() {
	var (
		response utils.Response
		paramURL = "5f24d45125ea51bc11111111"
	)

	// Setup request
	url := "/companies" + paramURL + "transactions"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// Run HTTP server
	suite.e.ServeHTTP(rec, req)

	// Parse
	json.Unmarshal([]byte(rec.Body.String()), &response)

	// Test
	assert.Equal(suite.T(), http.StatusNotFound, rec.Code)
	assert.Equal(suite.T(), nil, response["data"])
}

func TestCompanySuite(t *testing.T) {
	suite.Run(t, new(CompanyCreateSuite))
	suite.Run(t, new(TransactionFindByCompanyIDSuite))
}

func removeOldDataCompany() {
	database.CompanyCol().DeleteMany(context.Background(), bson.M{})
}
