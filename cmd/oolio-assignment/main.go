package main

import (
	"log"
	"oolio-assignment/pkg/handler"
	"oolio-assignment/pkg/repository"
)

func main() {
	var processConfig handler.ProcessConfig
	db, err := repository.SetUpDB()
	if err != nil {
		log.Fatal("failed to setup setup Database %w", err)
	}

	CurdRepo := repository.NewCurdRepo(db)
	processConfig.CurdRepo = CurdRepo
	MigrateDBRepo(db)

	log.Println("Please wait loading valid coupons process started")
	validCoupons, err := LoadValidCoupons()
	if err != nil {
		log.Fatalf("Failed to load valid coupons: %v", err)
	}
	log.Println("loading coupons process complete")

	initWebServer(processConfig, validCoupons)
}
