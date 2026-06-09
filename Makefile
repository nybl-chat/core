APP_NAME := nybl
GO := go

.PHONY: data auth dev

gen:
	@protoc \
		pkg/proto/struct/*.proto \
		--go_out=pkg/proto/gen
		
		
# ! ONLY FOR DEVELOPMENT : DO NOT USE IN PRODUCTION
dev:
	bash build.sh
	@trap 'kill 0' EXIT; \
		./bin/data & \
		./bin/auth & \
		wait