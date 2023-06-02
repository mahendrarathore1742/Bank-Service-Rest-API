build:
	@go build -o bin/bank-service-rest-api

run: build
	@./bin/bank-service-rest-api

test:
	@go test -v ./...