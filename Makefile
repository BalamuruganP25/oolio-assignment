SERVICE_NAME = oolio-assignment

export GO111MODULE = on

run: build
	@docker compose up
	# @docker run -d -p 8089:8089 --name $(SERVICE_NAME) $(SERVICE_NAME)

build:
	@docker compose build
	# @docker build -t $(SERVICE_NAME) .

stop:
	@docker compose down
	# @docker stop $(SERVICE_NAME)
	# @docker rm $(SERVICE_NAME)

dep:
	@go mod tidy
	@go mod vendor
