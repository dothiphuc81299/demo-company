package test

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
	"demo-company/apptest"
)


type CompanyCreateSuite struct {
	suite.Suite
	e *echo.Echo
	data models.CompanyCreatePayload
}

func (suite *CompanyCreateSuite) SetupSuite() {
	// init server ... 
	suite.e= apptest.InitServer()

	// clear data
	RemoveOldDataCompany()

	// set up payload data 
	suite.data = models.CompanyCreatePayload{
		Name :"Phuc Mars",
		CashbackPercent:19.2,
	}
}

func (suite *CompanyCreateSuite) TearDownSuite() {
	RemoveOldDataCompany()
}

func RemoveOldDataCompany() {
	database.CompanyCol().DeleteMany(context.Background(), bson.M{})
}

func (suite *CompanyCreateSuite) TestCompanyCreateSuccess() {
	var (
			response util.Response
			payload = suite.data	
	)
	
	//set up request
	req, _ := http.NewRequest(http.MethodPost, "/companies", util.HelperToIOReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	
	 // Run HTTP server
	suite.e.ServeHTTP(res, req)

	//parse .. 
	json.Unmarshal([]byte(res.Body.String()), &response)	
	
	//Test
	assert.Equal(suite.T(), http.StatusOK, res.Code)
	assert.NotEqual(suite.T(), nil, response["data"])
	assert.Equal(suite.T(), "thanh cong!", response["message"])

}

func (suite *CompanyCreateSuite) TestCompanyCreateFailureWithInvalidName() {
		var (
			response util.Response
			payload = models.CompanyCreatePayload{
				Name:           "",
				CashbackPercent: 19.2,
			}
	)
	
	//set up request
	req, _ := http.NewRequest(http.MethodPost, "/companies", util.HelperToIOReader(payload))
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

func (suite *CompanyCreateSuite) TestCompanyCreateFailureWithCashbackPercent() {
		var (
			response util.Response
			payload = models.CompanyCreatePayload{
				Name:           "PhucMassr",
				CashbackPercent: 0,
			}
	)

	//set up request
	req, _ := http.NewRequest(http.MethodPost, "/companies", util.HelperToIOReader(payload))
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

type TransactionFindByCompanyIDSuite struct {
	suite.Suite
	e        *echo.Echo
	paramURL string
}

func (suite *TransactionFindByCompanyIDSuite) SetupSuite() {
	// set up server ... 
	suite.e =apptest.InitServer()

	// clear data ...
	RemoveOldDataCompany()

	//set up data 
	suite.paramURL = util.HelperCompanyCreateFake()
}

func (suite *TransactionFindByCompanyIDSuite) TearDownSuite() {
	RemoveOldDataCompany()
}


func (suite *TransactionFindByCompanyIDSuite) TestTransactionFindByCompanyIDSuccess() {
	var (
		response util.Response
	)

	// Setup request
	url := "/companies/" + suite.paramURL + "/transactions"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// Run HTTP server
	suite.e.ServeHTTP(rec, req)

	//Parse
	json.Unmarshal([]byte(rec.Body.String()), &response)

	//Test
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	assert.Equal(suite.T(), "thanh cong!", response["message"])
}

func (suite *TransactionFindByCompanyIDSuite) TestTransactionFindByCompanyIDFailureWithInvalidUCompanyID() {
	var (
		response util.Response
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
		response util.Response
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
