# Variables
SERVICE_IMAGE = url-shortener-service

# Build the service Docker image
build:
	docker build -t $(SERVICE_IMAGE) .

# Run the service Docker container
run:
	docker run -d -p 8080:8080 --name $(SERVICE_IMAGE) $(SERVICE_IMAGE)

# Stop the service Docker container
stop:
	docker stop $(SERVICE_IMAGE) || true
	docker rm $(SERVICE_IMAGE) || true

# Build and run the service
up: build run

# Stop the service
down: stop

# Rebuild and run the service
restart: down up

.PHONY: build run stop up down restart
