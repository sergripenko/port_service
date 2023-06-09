# Path to docker-compose file
PATH_DOCKER_COMPOSE_FILE=$(realpath ./deploy/docker-compose.yml)
# Docker compose starting options.
DOCKER_COMPOSE_OPTIONS= -f $(PATH_DOCKER_COMPOSE_FILE)

golangci: ## Linters
	golangci-lint run -v

test: ## Testing
	go test ./internal/service/... -v

build:
	docker build . -t ports-service

run:
	docker run -it ports-service

service-build: ## Build service and all it's dependencies
	docker compose $(DOCKER_COMPOSE_OPTIONS) build --no-cache

service-start-dependencies: ## Start service dependencies in Docker
	@echo ">>> Start all Service dependencies."
	docker compose $(DOCKER_COMPOSE_OPTIONS) up

service-start: service_build service-start-dependencies ## Start service in Docker
	@echo ">>> Starting service."
	@echo ">>> Starting up service container."
	docker compose $(DOCKER_COMPOSE_OPTIONS) up $(SERVICE)

service-stop: ## Stop service in Docker
	@echo ">>> Stopping service."
	@docker compose $(DOCKER_COMPOSE_OPTIONS) down -v --remove-orphans

service-restart: service-stop service-start ## Restart service in Docker

clean-mocks: ## Cleans old mocks
	find . -name "mock_*.go" -type f -print0 | xargs -0 /bin/rm -f

mocks: clean-mocks ## For generating mock based on all project interfaces
	mockery --all --dir "./internal/service/port" --inpackage --case underscore