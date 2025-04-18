package order

import (
	"errors"
)

func validateOrderReq(req orderRequest, validCoupons map[string]bool) error {

	if len(req.Items) == 0 {
		return errors.New("order iteams should not be empty")
	}
	if req.CouponCode != "" && !validCoupons[req.CouponCode] {
		return errors.New("invalid coupon code")
	}

	return nil
}
