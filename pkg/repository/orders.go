package repository

import (
	"context"
)

func (repo *CurdRepository) CreateOrder(ctx context.Context, orderID, couponCode string, orders []byte, products []byte) error {
	query := `
        INSERT INTO orders (id, coupon_code, orders, products)
        VALUES ($1, $2, $3, $4)
    `
	_, err := repo.db.ExecContext(ctx, query, orderID, couponCode, orders, products)
	if err != nil {
		return err
	}
	return nil
}
