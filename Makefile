golangci: ## Linters
	golangci-lint run -v

test: ## Testing
	go test ./internal/service/... -v

build:
	docker build . -t ports-service

run:
	docker run -it ports-service