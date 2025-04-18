package product_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"oolio-assignment/pkg/handler"
	"oolio-assignment/pkg/handler/product"
	"oolio-assignment/pkg/mocks"
	"oolio-assignment/pkg/repository"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestCreatProduct(t *testing.T) {
	suite.Run(t, new(creatProductTestSuite))
}

type creatProductTestSuite struct {
	suite.Suite
	router   *chi.Mux
	recorder *httptest.ResponseRecorder
	CurdRepo *mocks.CrudRepo
}

func (s *creatProductTestSuite) SetupTest() {

	s.recorder = httptest.NewRecorder()
	s.router = chi.NewRouter()
	s.CurdRepo = new(mocks.CrudRepo)
	config := handler.ProcessConfig{
		CurdRepo: s.CurdRepo,
	}
	s.router.Post("/v1/product", product.CreateProduct(&config))
}

func (s *creatProductTestSuite) executeTestSuiteRequest(reqbody string) {
	req := httptest.NewRequest("POST", "/v1/product", strings.NewReader(reqbody))
	s.router.ServeHTTP(s.recorder, req)

}
func (s *creatProductTestSuite) TestCreateProductSuccess() {
	s.CurdRepo.On("CreateProduct", mock.Anything, mock.MatchedBy(
		func(dbobj *repository.ProductModel) bool {
			s.NotNil(dbobj)

			return true
		},
	)).Return("10", nil)

	req := fmt.Sprintf(`{
	  "name":"%s",
	  "price":%f,
	  "category":"%s"

	}`, "butter", 10.5, "veg")
	s.executeTestSuiteRequest(req)
	s.Equal(http.StatusCreated, s.recorder.Code)
	s.JSONEq(`{"category":"veg", "id":"10", "name":"butter", "price":10.5}`, s.recorder.Body.String())

}

func (s *creatProductTestSuite) TestCreateProductFailedToCreateProduct() {
	s.CurdRepo.On("CreateProduct", mock.Anything, mock.Anything).Return("", errors.New("DB error"))
	req := fmt.Sprintf(`{
	  "name":"%s",
	  "price":%f,
	  "category":"%s"

	}`, "butter", 11.5, "veg")
	s.executeTestSuiteRequest(req)
	s.Equal(http.StatusInternalServerError, s.recorder.Code)
	s.JSONEq(`{
		"title":"internal server error",
		"details":"failed to create product - DB error"
		}`, s.recorder.Body.String())

}
func (s *creatProductTestSuite) TestCreateProductFailedProductAlreadExist() {
	s.CurdRepo.On("CreateProduct", mock.Anything, mock.Anything).Return("", errors.New("product already exists"))
	req := fmt.Sprintf(`{
	  "name":"%s",
	  "price":%f,
	  "category":"%s"

	}`, "butter", 11.5, "veg")
	s.executeTestSuiteRequest(req)
	s.Equal(http.StatusConflict, s.recorder.Code)
	s.JSONEq(`{"title":"conflict","details":"product already exists"}`, s.recorder.Body.String())

}

func (s *creatProductTestSuite) TestValidateProductReq() {
	tests := []struct {
		name       string
		reqBody    string
		wantStatus int
		wantError  string
	}{
		{
			name: "Missing Name",
			reqBody: `{
				"id": "1",
				"name": "",
				"price": 10.5,
				"category": "Fast Food"
			}`,
			wantStatus: http.StatusBadRequest,
			wantError:  "name shouldn't be empty",
		},
		{
			name: "Missing Price",
			reqBody: `{
				"id": "1",
				"name": "Burger",
				"price": 0,
				"category": "Fast Food"
			}`,
			wantStatus: http.StatusBadRequest,
			wantError:  "price shouldn't be empty",
		},
	}

	for _, tc := range tests {
		s.T().Run(tc.name, func(t *testing.T) {
			s.executeTestSuiteRequest(tc.reqBody)
			s.Equal(tc.wantStatus, s.recorder.Code)
			var errRes handler.ErrResponse
			err := json.NewDecoder(s.recorder.Body).Decode(&errRes)
			if err != nil {
				log.Fatal("failed to run the test case")
			}
			s.Equal(tc.wantError, errRes.Details)
		})
	}
}
