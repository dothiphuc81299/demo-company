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


type CompanySuite struct {
	suite.Suite
	e *echo.Echo
}

func (suite CompanySuite) SetupSuite() {
	suite.e= InitEcho()
	RemoveOldDataCompany()
}

func (suite CompanySuite) TearDownSuite() {
	RemoveOldDataCompany()
}

func RemoveOldDataCompany() {
	database.CompanyCol().DeleteMany(context.Background(), bson.M{})
}

func (suite *CompanySuite) TestSuccess() {
	var (
			response util.Response
		payload = models.CompanyCreatePayload{
			Name:           "PhucMars",
			CashbackPercent: 19.2,
		}
	)
	suite.e = InitEcho()

	//set up request
	req, _ := http.NewRequest(http.MethodPost, "/companies", util.HelperToIOReader(payload))
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

func (suite *CompanySuite) TestNameFail() {
		var (
			response util.Response
		payload = models.CompanyCreatePayload{
			Name:           "",
			CashbackPercent: 19.2,
		}
	)
	suite.e = InitEcho()

	//set up request
	req, _ := http.NewRequest(http.MethodPost, "/companies", util.HelperToIOReader(payload))
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


func (suite *CompanySuite) TestCashbackFail() {
		var (
			response util.Response
		payload = models.CompanyCreatePayload{
			Name:           "PhucMassr",
			CashbackPercent: 0,
		}
	)
	suite.e = InitEcho()

	//set up request
	req, _ := http.NewRequest(http.MethodPost, "/companies", util.HelperToIOReader(payload))
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
func TestCompanySuite(t *testing.T) {
	suite.Run(t, new(CompanySuite))
}
