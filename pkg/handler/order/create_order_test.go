package order_test

import (
	"net/http"
	"net/http/httptest"
	"oolio-assignment/pkg/handler"
	"oolio-assignment/pkg/handler/order"
	"oolio-assignment/pkg/mocks"
	"oolio-assignment/pkg/repository"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestCreatOrder(t *testing.T) {
	suite.Run(t, new(creatOrderTestSuite))
}

type creatOrderTestSuite struct {
	suite.Suite
	router   *chi.Mux
	recorder *httptest.ResponseRecorder
	CurdRepo *mocks.CrudRepo
}

func (s *creatOrderTestSuite) SetupTest() {

	s.recorder = httptest.NewRecorder()
	s.router = chi.NewRouter()
	s.CurdRepo = new(mocks.CrudRepo)
	config := handler.ProcessConfig{
		CurdRepo: s.CurdRepo,
	}
	validCoupons := map[string]bool{
		"FIFTYOFF": true,
	}
	s.router.Post("/v1/order", order.CreateOrder(&config, validCoupons))
}

func (s *creatOrderTestSuite) executeTestSuiteRequest(reqbody string) {
	req := httptest.NewRequest("POST", "/v1/order", strings.NewReader(reqbody))
	s.router.ServeHTTP(s.recorder, req)

}

func (s *creatOrderTestSuite) TestCreateOrderSuccess_withCouponCode() {
	s.CurdRepo.On("CreateOrder", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	s.CurdRepo.On("GetProductByProductID", mock.Anything, 120).Return(repository.ProductModel{
		ID:       "120",
		Name:     "Butter",
		Price:    10.9,
		Category: "Dairy",
	}, nil)

	reqBody := `{
		"coupon_code": "FIFTYOFF",
		"items": [
			{ "product_id": "120", "quantity": 2 }
		]
	}`
	s.executeTestSuiteRequest(reqBody)
	s.Equal(http.StatusCreated, s.recorder.Code)
	s.JSONEq(`{
		"coupon_code":"FIFTYOFF",
		"items":[{"product_id":"120","quantity":2}],
		"products":[{"id":"120","name":"Butter","price":10.9,"category":"Dairy"}]}`, s.recorder.Body.String())

}

func (s *creatOrderTestSuite) TestCreateOrderFailed_InvalidCouponCode() {
	reqBody := `{
		"coupon_code": "FIFT893",
		"items": [
			{ "product_id": "120", "quantity": 2 }
		]
	}`
	s.executeTestSuiteRequest(reqBody)
	s.Equal(http.StatusBadRequest, s.recorder.Code)
	s.JSONEq(`{"tittle":"validation error","details":"invalid coupon code"}`, s.recorder.Body.String())

}

func (s *creatOrderTestSuite) TestCreateOrderFailed_InvalidProductID() {

	s.CurdRepo.On("GetProductByProductID", mock.Anything, 120).Return(repository.ProductModel{},
		repository.RecordNotFound)

	reqBody := `{
		"coupon_code": "FIFTYOFF",
		"items": [
			{ "product_id": "120", "quantity": 2 }
		]
	}`
	s.executeTestSuiteRequest(reqBody)
	s.Equal(http.StatusBadRequest, s.recorder.Code)
	s.JSONEq(`{"tittle":"invalid product id","details":"invaild product id : 120"}`, s.recorder.Body.String())

}

func (s *creatOrderTestSuite) TestCreateOrderFailed_InvalidOrder() {
	reqBody := `{
		"coupon_code": "FIFTYOFF",
		"items": []
	}`
	s.executeTestSuiteRequest(reqBody)
	s.Equal(http.StatusBadRequest, s.recorder.Code)
	s.JSONEq(`{"tittle":"validation error","details":"order iteams should not be empty"}`, s.recorder.Body.String())

}
