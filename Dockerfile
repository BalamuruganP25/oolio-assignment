FROM golang:1.24.1-alpine3.21 AS builder-env

WORKDIR /go/src/

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o app /go/src/cmd/oolio-assignment
RUN cp app /app
EXPOSE 8080
CMD ["/app"]
