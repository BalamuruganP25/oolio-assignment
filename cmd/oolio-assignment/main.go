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
	initWebServer(processConfig)
}
