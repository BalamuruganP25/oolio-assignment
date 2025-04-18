package repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"oolio-assignment/pkg/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type productRepoSuite struct {
	suite.Suite
	db       *sql.DB
	mock     sqlmock.Sqlmock
	curdRepo *repository.CurdRepository
}

func TestproductRepoSuite(t *testing.T) {
	suite.Run(t, new(productRepoSuite))
}

func (d *productRepoSuite) SetupTest() {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	d.Require().NoError(err)
	d.mock = mock
	d.db = db
	d.curdRepo = repository.NewCurdRepo(d.db)
}

func (d *productRepoSuite) TestCreateProduct_Success() {
	product := &repository.ProductModel{
		Name:     "Toothpaste",
		Price:    50,
		Category: "Health",
	}
	expectedID := "123"

	d.mock.ExpectQuery("INSERT INTO products").
		WithArgs(product.Name, product.Price, product.Category).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedID))

	id, err := d.curdRepo.CreateProduct(context.TODO(), product)
	d.NoError(err)
	d.Equal(expectedID, id)
	d.NoError(d.mock.ExpectationsWereMet())
}

func (d *productRepoSuite) TestCreateProduct_DuplicateKey() {
	product := &repository.ProductModel{
		Name:     "Soap",
		Price:    30,
		Category: "Personal",
	}

	d.mock.ExpectQuery("INSERT INTO products").
		WithArgs(product.Name, product.Price, product.Category).
		WillReturnError(fmt.Errorf("pq: duplicate key value violates unique constraint"))

	id, err := d.curdRepo.CreateProduct(context.TODO(), product)
	d.ErrorIs(err, repository.ProductExist)
	d.Empty(id)
	d.NoError(d.mock.ExpectationsWereMet())
}

func (d *productRepoSuite) TestGetProductByProductID_Success() {
	expected := repository.ProductModel{
		ID:       "1",
		Name:     "Shampoo",
		Price:    100,
		Category: "Haircare",
	}

	d.mock.ExpectQuery("SELECT id, name, price, category FROM products WHERE id = \\$1").
		WithArgs(expected.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price", "category"}).
			AddRow(expected.ID, expected.Name, expected.Price, expected.Category))

	product, err := d.curdRepo.GetProductByProductID(context.TODO(), 1)
	d.NoError(err)
	d.Equal(expected, product)
	d.NoError(d.mock.ExpectationsWereMet())
}

func (d *productRepoSuite) TestGetProductByProductID_NotFound() {
	productID := 999

	d.mock.ExpectQuery("SELECT id, name, price, category FROM products WHERE id = \\$1").
		WithArgs(productID).
		WillReturnError(sql.ErrNoRows)

	product, err := d.curdRepo.GetProductByProductID(context.TODO(), productID)
	d.ErrorIs(err, repository.RecordNotFound)
	d.Equal(repository.ProductModel{}, product)
	d.NoError(d.mock.ExpectationsWereMet())
}

func (d *productRepoSuite) TestGetProduct_Success() {
	expected := []repository.ProductModel{
		{ID: "1", Name: "Toothbrush", Price: 25, Category: "Dental"},
		{ID: "2", Name: "Shampoo", Price: 70, Category: "Haircare"},
	}

	d.mock.ExpectQuery("SELECT id, name, price, category FROM products LIMIT 2 OFFSET 0").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price", "category"}).
			AddRow(expected[0].ID, expected[0].Name, expected[0].Price, expected[0].Category).
			AddRow(expected[1].ID, expected[1].Name, expected[1].Price, expected[1].Category))

	products, err := d.curdRepo.GetProduct(context.TODO(), "1", "2")
	d.NoError(err)
	d.Equal(expected, products)
	d.NoError(d.mock.ExpectationsWereMet())
}
