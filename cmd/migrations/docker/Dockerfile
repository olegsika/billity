FROM golang:1.15 as migrations-builder

# Update go modules
RUN go mod tidy
RUN go mod vendor

RUN	go build -o ./package/migrations/migrations ./cmd/migrations/main.go