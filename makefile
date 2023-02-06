TAG := teas/backend:latest

build:
	@echo "Building Backend docker image"
	@docker build -f deploy/docker/Dockerfile -t $(TAG) .

start:
	@echo "Starting Backend..."
	@sh ./deploy/scripts/up.sh
	@echo "Generating Swagger"
	swag init

start-dev:
	@echo "Starting Backend..."
	@sh ./deploy/scripts/up-dev.sh
	@echo "Generating Swagger"
	swag init
		
stop:
	@echo "Stopping Backend..."
	@sh ./deploy/scripts/down.sh

stop-dev:
	@echo "Stopping Backend..."
	@sh ./deploy/scripts/down-dev.sh
	
.PHONY: build start stop start-dev stop-dev