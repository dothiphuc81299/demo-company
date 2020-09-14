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


type CompanySuite struct {
	suite.Suite
	e *echo.Echo
	data models.CompanyCreatePayload
}

func (suite *CompanySuite) SetupSuite() {
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

func (suite *CompanySuite) TearDownSuite() {
	RemoveOldDataCompany()
}

func RemoveOldDataCompany() {
	database.CompanyCol().DeleteMany(context.Background(), bson.M{})
}

func (suite *CompanySuite) TestCompanyCreateSuccess() {
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

func (suite *CompanySuite) TestCompanyCreateFailureWithInvalidName() {
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

func (suite *CompanySuite) TestCompanyCreateFailureWithCashbackPercent() {
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
func TestCompanySuite(t *testing.T) {
	suite.Run(t, new(CompanySuite))
}
