package product

import (
	"fmt"
	"net/http"
	"oolio-assignment/pkg/handler"
	"oolio-assignment/pkg/repository"
	"strconv"

	"github.com/go-chi/chi"
)

func GetProductByID(config *handler.ProcessConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productIDStr := chi.URLParam(r, "product_id")
		if productIDStr == "" {
			handler.ErrorResponse(w, http.StatusBadRequest, handler.ErrResponse{
				Tittle:  "validation error",
				Details: "product id should be empty",
			})
			return
		}

		product_id, err := strconv.Atoi(productIDStr)
		if err != nil {
			handler.ErrorResponse(w, http.StatusBadRequest, handler.ErrResponse{
				Tittle:  "invalid Request",
				Details: "product id should not be empty",
			})

			return
		}
		
		// Get product by product id 
		product, err := config.CurdRepo.GetProductByProductID(r.Context(), product_id)
		if err != nil {
			if err == repository.RecordNotFound {
				handler.ErrorResponse(w, http.StatusNotFound, handler.ErrResponse{
					Tittle:  "Resource Not Found",
					Details: "product not found",
				},
				)
				return
			}
			handler.ErrorResponse(w, http.StatusInternalServerError,
				handler.ErrResponse{
					Tittle:  "internal server error",
					Details: fmt.Errorf("failed to get product - %w", err).Error(),
				},
			)
			return
		}

		// sent response
		handler.SendResponse(w, product, http.StatusOK)

	}
}
