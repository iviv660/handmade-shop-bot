.PHONY: test up down
test:
	@go test ./internal/service/tests

up:
	@docker compose up -d

down:
	@docker compose down