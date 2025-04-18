package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	"oolio-assignment/pkg/handler"
	"oolio-assignment/pkg/repository"
)

func CreateProduct(s *handler.ProcessConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Product

		// Decode the request body
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {

			handler.ErrorResponse(w, http.StatusInternalServerError,
				handler.ErrResponse{
					Tittle:  "payload error",
					Details: fmt.Sprintf("invalid request : %v", err),
				},
			)
			return
		}

		// validate api request
		err = validateProductReq(req)
		if err != nil {
			handler.ErrorResponse(w, http.StatusBadRequest, handler.ErrResponse{
				Tittle:  "validation error",
				Details: err.Error(),
			},
			)
			return
		}

		modle := &repository.ProductModel{
			Name:     req.Name,
			Price:    req.Price,
			Category: req.Category,
		}

		// store product deatils
		modle.ID, err = s.CurdRepo.CreateProduct(r.Context(), modle)
		if err != nil {
			if err.Error() == repository.ProductExist.Error() {
				handler.ErrorResponse(w, http.StatusConflict,
					handler.ErrResponse{
						Tittle:  "conflict",
						Details: repository.ProductExist.Error(),
					},
				)
				return
			}

			handler.ErrorResponse(w, http.StatusInternalServerError,
				handler.ErrResponse{
					Tittle:  "internal server error",
					Details: fmt.Errorf("failed to create product - %w", err).Error(),
				},
			)
			return
		}

		// sent response
		handler.SendResponse(w, modle, http.StatusCreated)

	}
}
