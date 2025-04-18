package product

import (
	"errors"
)

func validateProductReq(req Product) error {

	if req.Name == "" {
		return errors.New("name shouldn't be empty")
	}

	if req.Price == 0 {
		return errors.New("price shouldn't be empty")
	}

	if req.Category == "" {
		return errors.New("category shouldn't be empty")
	}

	return nil
}
