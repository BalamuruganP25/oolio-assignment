package order

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"oolio-assignment/pkg/handler"
	"oolio-assignment/pkg/repository"
	"strconv"

	"github.com/google/uuid"
)

func CreateOrder(config *handler.ProcessConfig, validCoupons map[string]bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req orderRequest
		ctx := r.Context()

		// Decode the request body
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			handler.ErrorResponse(w, http.StatusBadRequest,
				handler.ErrResponse{
					Title:   "invalid payload",
					Details: fmt.Sprintf("invalid request : %v", err),
				},
			)
			return
		}

		// validate api request
		err = validateOrderReq(req, validCoupons)
		if err != nil {
			handler.ErrorResponse(w, http.StatusBadRequest, handler.ErrResponse{
				Title:   "validation error",
				Details: err.Error(),
			},
			)
			return
		}

		// get user order product list
		productList, err := getProductList(ctx, config, req.Items)
		if err != nil {
			handler.ErrorResponse(w, http.StatusBadRequest,
				handler.ErrResponse{
					Title:   "invalid product id",
					Details: err.Error(),
				},
			)
			return
		}

		products, _ := json.Marshal(productList)
		orders, _ := json.Marshal(req.Items)
		order_id := uuid.New().String()

		// create order request
		err = config.CurdRepo.CreateOrder(ctx, order_id, req.CouponCode, orders, products)
		if err != nil {
			handler.ErrorResponse(w, http.StatusInternalServerError,
				handler.ErrResponse{
					Title:   "internal server error",
					Details: fmt.Errorf("failed to create product - %w", err).Error(),
				},
			)
		}

		resp := repository.OrderModel{
			Id:         order_id,
			CouponCode: req.CouponCode,
			Items:      req.Items,
			Products:   productList,
		}

		// sent response
		handler.SendResponse(w, resp, http.StatusCreated)

	}
}

func getProductList(ctx context.Context, cs *handler.ProcessConfig,
	orderIteam []repository.OrderIteamModel) ([]repository.ProductModel, error) {
	var products []repository.ProductModel

	for i := 0; i < len(orderIteam); i++ {

		productID, err := strconv.Atoi(orderIteam[i].ProductId)
		if err != nil {
			return []repository.ProductModel{}, fmt.Errorf("invalid product id : %s", orderIteam[i].ProductId)
		}
		product, err := cs.CurdRepo.GetProductByProductID(ctx, productID)
		if err != nil {
			if err == repository.RecordNotFound {
				return []repository.ProductModel{}, fmt.Errorf("invaild product id : %s", orderIteam[i].ProductId)
			}
			return []repository.ProductModel{}, fmt.Errorf("failed to get product id : %w", err)
		}
		products = append(products, product)
	}
	return products, nil
}
