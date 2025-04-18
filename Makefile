# 🛠️ Makefile for Oolio Assignment Project

SERVICE_NAME = oolio-assignment

export GO111MODULE = on

.PHONY: run build stop dep test

## 🔄 Run the full Docker stack
run: build
	@echo "🚀 Starting $(SERVICE_NAME)..."
	@docker compose up

## 🏗️ Build Docker containers
build:
	@echo "🔧 Building Docker images..."
	@docker compose build

## 🛑 Stop and remove containers
stop:
	@echo "🧹 Stopping services..."
	@docker compose down

## 📦 Install and tidy Go dependencies
dep:
	@echo "📦 Tidying and vendoring Go dependencies..."
	@go mod tidy
	@go mod vendor

## 🧪 Run all Go tests
test:
	@echo "🧪 Running tests..."
	@go test -v ./...

## 🔍 Run code linter
lint:
	@echo "🔍 Running golangci-lint..."
	@golangci-lint run ./...