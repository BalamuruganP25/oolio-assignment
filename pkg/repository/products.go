package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

type CurdRepository struct {
	db *sql.DB
}

func NewCurdRepo(db *sql.DB) *CurdRepository {
	return &CurdRepository{
		db: db,
	}
}

func (repo *CurdRepository) CreateProduct(ctx context.Context, req *ProductModel) (string, error) {
	var id string
	query := `
        INSERT INTO products (name, price, category)
        VALUES ($1, $2, $3)
        RETURNING id::text
    `
	err := repo.db.QueryRow(query, req.Name, req.Price, req.Category).Scan(&id)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return "", ProductExist
		}
		return "", err
	}

	return id, nil

}

func (repo *CurdRepository) GetProductByProductID(ctx context.Context, productID int) (ProductModel, error) {
	query := `SELECT id, name, price, category FROM products WHERE id = $1`
	row := repo.db.QueryRow(query, productID)

	var product ProductModel
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Category)
	if err != nil {
		if err == sql.ErrNoRows {
			return ProductModel{}, RecordNotFound
		}
		return ProductModel{}, err
	}

	return product, nil

}

func (repo *CurdRepository) GetProduct(ctx context.Context, page, limit string) ([]ProductModel, error) {
	query := buildGetListProductQuery(page, limit)
	rows, err := repo.db.Query(query)
	if err != nil {
		return []ProductModel{}, err
	}
	var product []ProductModel
	for rows.Next() {
		var s ProductModel
		err := rows.Scan(&s.ID, &s.Name, &s.Price, &s.Category)
		if err != nil {
			return []ProductModel{}, err
		}
		product = append(product, s)
	}

	return product, nil

}

func buildGetListProductQuery(pageStr, limitStr string) string {
	query := `SELECT id, name, price, category FROM products LIMIT %d OFFSET %d `
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	query = fmt.Sprintf(query, limit, offset)
	return query
}
