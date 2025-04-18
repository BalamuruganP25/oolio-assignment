package order

import "oolio-assignment/pkg/repository"

type orderRequest struct {
	CouponCode string                       `json:"coupon_code"`
	Items      []repository.OrderIteamModel `json:"items"`
}
