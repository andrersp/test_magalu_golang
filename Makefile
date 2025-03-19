.PHONY: dev
dev:
	air

.PHONY: test
test:
	go test -v -cover ./...

.PHONY: mod
mod:
	go mod download
	go mod tidy	

.PHONY: swag
swag:
	swag init -g ./cmd/api/main.go