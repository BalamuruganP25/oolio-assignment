package product_test

import (
	"encoding/json"
	"errors"
	"fmt"
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

func TestGetProductByProductID(t *testing.T) {
	suite.Run(t, new(getProductByProductIDTestSuite))
}

type getProductByProductIDTestSuite struct {
	suite.Suite
	router   *chi.Mux
	recorder *httptest.ResponseRecorder
	CurdRepo *mocks.CrudRepo
}

func (s *getProductByProductIDTestSuite) SetupTest() {
	s.recorder = httptest.NewRecorder()
	s.router = chi.NewRouter()
	s.CurdRepo = new(mocks.CrudRepo)

	config := handler.ProcessConfig{
		CurdRepo: s.CurdRepo,
	}
	s.router.Get("/v1/product/{product_id}", product.GetProductByID(&config))
}

func (s *getProductByProductIDTestSuite) executeRequest(productID string) {
	req := httptest.NewRequest("GET", fmt.Sprintf("/v1/product/%s", productID), nil)
	s.router.ServeHTTP(s.recorder, req)
}

func (s *getProductByProductIDTestSuite) TestGetProductByProductIDSuccess() {
	expectedProduct := repository.ProductModel{
		ID:       "123",
		Name:     "Test Product",
		Price:    49.99,
		Category: "Books",
	}
	s.CurdRepo.On("GetProductByProductID", mock.Anything, 123).Return(expectedProduct, nil)
	s.executeRequest("123")
	var actualProduct repository.ProductModel
	err := json.NewDecoder(s.recorder.Body).Decode(&actualProduct)
	assert.NoError(s.T(), err)
	s.Equal(expectedProduct, actualProduct)
	s.Equal(http.StatusOK, s.recorder.Code)
	s.CurdRepo.AssertExpectations(s.T())
}

func (s *getProductByProductIDTestSuite) TestGetProductByProductIDFailed_Returns404() {
	s.CurdRepo.On("GetProductByProductID", mock.Anything, 10).Return(repository.ProductModel{}, repository.RecordNotFound)
	s.executeRequest("10")
	s.Equal(http.StatusNotFound, s.recorder.Code)
	s.CurdRepo.AssertExpectations(s.T())
}

func (s *getProductByProductIDTestSuite) TestGetProductByProductIDFailed_Returns500() {
	s.CurdRepo.On("GetProductByProductID", mock.Anything, 123).Return(repository.ProductModel{}, errors.New("unexpected DB error"))
	s.executeRequest("123")
	s.Equal(http.StatusInternalServerError, s.recorder.Code)
	s.CurdRepo.AssertExpectations(s.T())
}
