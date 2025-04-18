package repository

import (
	"context"
	"errors"
)

type ProductModel struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

type OrderModel struct {
	Id         string            `json:"id"`
	CouponCode string            `json:"coupon_code"`
	Items      []OrderIteamModel `json:"items"`
	Products   []ProductModel    `json:"products"`
}

type OrderIteamModel struct {
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

// errors

var (
	ProductExist   = errors.New("product already exists")
	RecordNotFound = errors.New("record not found")
)

type CrudRepo interface {

	// product operation
	CreateProduct(ctx context.Context, req *ProductModel) (string, error)
	GetProductByProductID(ctx context.Context, productID int) (ProductModel, error)
	GetProduct(ctx context.Context, page, limit string) ([]ProductModel, error)

	// order operation
	CreateOrder(ctx context.Context, orderID, couponCode string, orders []byte, products []byte) error
}
