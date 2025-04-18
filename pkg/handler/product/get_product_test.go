package product_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"oolio-assignment/pkg/handler"
	"oolio-assignment/pkg/handler/product"
	"oolio-assignment/pkg/mocks"
	"oolio-assignment/pkg/repository"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestGetProduct(t *testing.T) {
	suite.Run(t, new(getProductTestSuite))
}

type getProductTestSuite struct {
	suite.Suite
	router   *chi.Mux
	recorder *httptest.ResponseRecorder
	CurdRepo *mocks.CrudRepo
}

func (s *getProductTestSuite) SetupTest() {
	s.recorder = httptest.NewRecorder()
	s.router = chi.NewRouter()
	s.CurdRepo = new(mocks.CrudRepo)

	config := handler.ProcessConfig{
		CurdRepo: s.CurdRepo,
	}
	s.router.Get("/v1/product", product.GetProduct(&config))
}

func (s *getProductTestSuite) executeGetProductTestSuiteRequest(url string) {
	req := httptest.NewRequest("GET", url, nil)
	s.router.ServeHTTP(s.recorder, req)

}

func (s *getProductTestSuite) TestGetProductListSuccess_WithOutFilter() {
	expectedProduct := []repository.ProductModel{
		{
			ID:       "123",
			Category: "veg",
			Name:     "butter",
			Price:    10.9,
		},
		{
			ID:       "121",
			Category: "non veg",
			Name:     "butter",
			Price:    10.9,
		},
		{
			ID:       "123",
			Category: "non veg",
			Name:     "butter",
			Price:    10.9,
		},
	}

	s.CurdRepo.On("GetProduct", mock.Anything, "", "").Return(expectedProduct, nil)

	requrl := "/v1/product"
	s.executeGetProductTestSuiteRequest(requrl)
	var actualProduct []repository.ProductModel
	err := json.NewDecoder(s.recorder.Body).Decode(&actualProduct)
	assert.NoError(s.T(), err)
	s.Equal(expectedProduct, actualProduct)
	s.Equal(http.StatusOK, s.recorder.Code)
	s.CurdRepo.AssertExpectations(s.T())
}

func (s *getProductTestSuite) TestGetProductSuccess_WithFilter() {
	expectedProduct := []repository.ProductModel{
		{
			ID:       "123",
			Category: "veg",
			Name:     "butter",
			Price:    10.9,
		},
	}

	s.CurdRepo.On("GetProduct", mock.Anything, "1", "10").Return(expectedProduct, nil)

	requrl := "/v1/product?page=1&limit=10"
	s.executeGetProductTestSuiteRequest(requrl)
	var actualProduct []repository.ProductModel
	err := json.NewDecoder(s.recorder.Body).Decode(&actualProduct)
	assert.NoError(s.T(), err)
	s.Equal(expectedProduct, actualProduct)
	s.Equal(http.StatusOK, s.recorder.Code)
	s.CurdRepo.AssertExpectations(s.T())
}

func (s *getProductTestSuite) TestGetProductFailed_DBError() {
	s.CurdRepo.On("GetProduct", mock.Anything, "1", "10").Return([]repository.ProductModel{}, errors.New("DB error"))
	requrl := "/v1/product?page=1&limit=10"
	s.executeGetProductTestSuiteRequest(requrl)
	s.Equal(http.StatusInternalServerError, s.recorder.Code)
	s.CurdRepo.AssertExpectations(s.T())
}
