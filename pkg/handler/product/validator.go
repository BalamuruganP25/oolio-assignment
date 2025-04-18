package product

import (
	"errors"
)

func validateProductReq(req Product) error {

	if req.Name == "" {
		return errors.New("name should be empty")
	}

	if req.Price == 0 {
		return errors.New("price should be empty")
	}

	if req.Category == "" {
		return errors.New("category should be empty")
	}

	return nil
}
