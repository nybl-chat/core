APP_NAME := nybl
GO := go

.PHONY: data auth dev

# ! ONLY FOR DEVELOPMENT : DO NOT USE IN PRODUCTION
dev:
	bash build.sh
	@trap 'kill 0' EXIT; \
		./bin/data & \
		./bin/auth & \
		wait