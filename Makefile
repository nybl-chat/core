APP_NAME := nybl
GO := go

.PHONY: data auth dev

build-for:
	@echo "Building $(SERVICE)..."
	go build -o bin/$(SERVICE) ./cmd/$(SERVICE)/main.go

data:
	 $(MAKE) build-for SERVICE=data

auth:
	 $(MAKE) build-for SERVICE=auth

# ! ONLY FOR DEVELOPMENT : DO NOT USE IN PRODUCTION
dev:
	bash dev-build.sh
	@trap 'kill 0' EXIT; \
		./bin/data & \
		./bin/auth & \
		wait