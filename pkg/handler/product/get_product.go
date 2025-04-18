package product

import (
	"fmt"
	"net/http"
	"oolio-assignment/pkg/handler"
	"oolio-assignment/pkg/repository"
)

func GetProduct(config *handler.ProcessConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		pageStr := r.URL.Query().Get("page")
		limitStr := r.URL.Query().Get("limit")

		// get product list
		product, err := config.CurdRepo.GetProduct(r.Context(), pageStr, limitStr)
		if err != nil {
			handler.ErrorResponse(w, http.StatusInternalServerError,
				handler.ErrResponse{
					Title:  "internal server error",
					Details: fmt.Errorf("failed to get product - %w", err).Error(),
				},
			)
			return

		}

		// To avoid Null Respone
		if len(product) == 0 {
			product = []repository.ProductModel{}
		}
		// sent response
		handler.SendResponse(w, product, http.StatusOK)

	}
}
