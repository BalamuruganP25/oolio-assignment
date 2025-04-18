package main

import (
	"bufio"
	"compress/gzip"
	"database/sql"
	"log"
	"net/http"
	"oolio-assignment/pkg/handler"
	"oolio-assignment/pkg/handler/order"
	"oolio-assignment/pkg/handler/product"
	"os"
	"strings"
	"unicode"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pressly/goose/v3"
	httpSwagger "github.com/swaggo/http-swagger"
)

// MigrateDBRepo - executes the database migration file
func MigrateDBRepo(db *sql.DB) {

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "/go/src/migrations"); err != nil {
		panic(err)
	}
}

func initWebServer(config handler.ProcessConfig, validCoupons map[string]bool) {
	log.Println("server is started")
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/v1", func(v1 chi.Router) {
		// Products
		v1.Post("/product", product.CreateProduct(&config))
		v1.Get("/product/{product_id}", product.GetProductByID(&config))
		v1.Get("/product", product.GetProduct(&config))

		// Orders
		v1.Post("/order", order.CreateOrder(&config, validCoupons))
	})

	// Serve Swagger
	r.Get("/swagger/openapi.yaml", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./swagger/openapi.yaml") // Path is relative to where app runs
	}))

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/openapi.yaml"),
	))
	http.ListenAndServe(":8089", r)
}

// LoadValidCoupons - read and store valid coupones
func LoadValidCoupons() (map[string]bool, error) {
	files := []string{
		"/go/src/coupones/couponbase1.gz",
		"/go/src/coupones/couponbase2.gz",
		"/go/src/coupones/couponbase3.gz",
	}
	counter := make(map[string]int)

	for _, file := range files {
		seen := make(map[string]bool)

		f, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		gr, err := gzip.NewReader(f)
		if err != nil {
			return nil, err
		}
		defer gr.Close()

		scanner := bufio.NewScanner(gr)
		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			word := strings.ToUpper(strings.TrimSpace(scanner.Text()))
			if len(word) >= 8 && len(word) <= 10 && isAlpha(word) && !seen[word] {
				counter[word]++
				seen[word] = true
			}
		}
	}

	result := make(map[string]bool)
	for code, count := range counter {
		if count >= 2 {
			result[code] = true
		}
	}

	return result, nil
}

func isAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
