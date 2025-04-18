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

- Docker
- Docker Compose
- Go 1.18+

---

### 🛠️ Development Commands

Use the included `Makefile` to run the project easily.

| Command         | Description                     |
|-----------------|---------------------------------|
| `make run`      | Build and start services        |
| `make build`    | Build Docker containers         |
| `make stop`     | Stop and remove containers      |
| `make dep`      | Install Go dependencies         |

---

### 🐳 Running the Service

```bash
make run

### Stop the API service

```bash
make run

### 📚 API Documentation
You can explore and interact with the API using the Swagger UI at:

➡️ http://localhost:8089/swagger/index.html

The OpenAPI specification file can also be found at:

📄 http://localhost:8089/swagger/openapi.yaml