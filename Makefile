swag:
	swag init -g cmd/main.go

lint:
	golangci-lint run

test:
	go test -v ./...
