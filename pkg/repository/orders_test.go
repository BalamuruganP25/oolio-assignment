package repository_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"oolio-assignment/pkg/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type orderRepoSuite struct {
	suite.Suite
	db       *sql.DB
	mock     sqlmock.Sqlmock
	curdRepo *repository.CurdRepository
}

func TestOrderRepoSuite(t *testing.T) {
	suite.Run(t, new(orderRepoSuite))
}

func (d *orderRepoSuite) SetupTest() {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	d.Require().NoError(err)
	d.mock = mock
	d.db = db
	d.curdRepo = repository.NewCurdRepo(d.db)
}

func (d *orderRepoSuite) TestCreateOrder_Success() {
	orderID := "order123"
	couponCode := "SAVE10"
	items := []repository.OrderIteamModel{
		{ProductId: "prod1", Quantity: 2},
		{ProductId: "prod2", Quantity: 1},
	}
	products := []repository.ProductModel{
		{ID: "prod1", Name: "Shampoo", Price: 100},
		{ID: "prod2", Name: "Soap", Price: 50},
	}

	ordersJSON, err := json.Marshal(items)
	d.Require().NoError(err)

	productsJSON, err := json.Marshal(products)
	d.Require().NoError(err)

	d.mock.ExpectExec("INSERT INTO orders").
		WithArgs(orderID, couponCode, ordersJSON, productsJSON).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = d.curdRepo.CreateOrder(context.TODO(), orderID, couponCode, ordersJSON, productsJSON)
	d.NoError(err)

	d.NoError(d.mock.ExpectationsWereMet())
}

func (d *orderRepoSuite) TestCreateOrder_DBError() {
	orderID := "order123"
	couponCode := "FAIL10"
	items := []repository.OrderIteamModel{{ProductId: "prodX", Quantity: 5}}
	products := []repository.ProductModel{}

	ordersJSON, err := json.Marshal(items)
	d.Require().NoError(err)

	productsJSON, err := json.Marshal(products)
	d.Require().NoError(err)

	d.mock.ExpectExec("INSERT INTO orders").
		WithArgs(orderID, couponCode, ordersJSON, productsJSON).
		WillReturnError(sql.ErrConnDone)

	err = d.curdRepo.CreateOrder(context.TODO(), orderID, couponCode, ordersJSON, productsJSON)
	d.Error(err)
	d.EqualError(err, sql.ErrConnDone.Error())

	d.NoError(d.mock.ExpectationsWereMet())
}
