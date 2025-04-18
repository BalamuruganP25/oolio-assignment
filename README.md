# 🛍️ Oolio Assignment – Product & Order API

This service provides a RESTful API for managing products and creating orders. Built with Go and containerized using Docker, it offers endpoints to create and retrieve products, as well as to place orders using those products.

---

## 📦 Features

- Create and list products  
- Retrieve a product by ID  
- Create an order with product items  
- Optional coupon code support for orders  
- JSON-based API  
- Swagger/OpenAPI 3.0.3 Spec  
- Docker + Makefile setup for easy development  

---

## 🚀 Getting Started

### 📁 Prerequisites

- [Docker](https://www.docker.com/)  
- [Docker Compose](https://docs.docker.com/compose/)  
- [Go 1.18+](https://go.dev/dl/)

---

### 🛠️ Development Commands

Use the included `Makefile` to run and manage the project easily:

| Command         | Description                            |
|-----------------|----------------------------------------|
| `make run`      | Build and start services               |
| `make build`    | Build Docker containers                |
| `make stop`     | Stop and remove containers             |
| `make dep`      | Install Go dependencies                |
| `make test`     | Run all Go unit tests                  |
| `make lint`     | Run code linter using `golangci-lint`  |

---

### ⏳ Note on Server Startup

> **Heads up:** The server may take a little extra time to start during the first run.  
> This is because it reads and loads coupon data from the following files:
>
> - `couponbase1.gz`  
> - `couponbase2.gz`  
> - `couponbase3.gz`
>
> This is a one-time initialization to prepare the coupon data for efficient access during order processing.

---

### 🐳 Running the Service

Start the service with:

make run

### 🛑 Stop the API Service
Shut everything down cleanly:

make stop

### 🧪 To run all Go test cases:

make test

### 📚 API Documentation

You can explore and interact with the API using the Swagger UI at:

➡️ http://localhost:8089/swagger/index.html

The OpenAPI specification file can also be found at:

📄 http://localhost:8089/swagger/openapi.yaml

