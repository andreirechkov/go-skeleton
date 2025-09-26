include .env
export $(shell sed -n 's/^\([A-Za-z0-9_]\+\)=.*/\1/p' .env)

DB_URL=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable

dev:
	air -c .air.toml

run:
	go run ./cmd/api

migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" down 1

migrate-reset:
	migrate -path migrations -database "$(DB_URL)" down

migrate-new:
	@if [ -z "$(name)" ]; then echo "Usage: make migrate-new name=feature_name"; exit 1; fi
	migrate create -dir migrations -ext sql -seq $(name)

migrate-force:
	@if [ -z "$(version)" ]; then echo "Usage: make migrate-force version=NUM"; exit 1; fi
	migrate -path migrations -database "$(DB_URL)" force $(version)

migrate-version:
	migrate -path migrations -database "$(DB_URL)" version || true

migrate-rm-last:
	@if [ -z "$(count)" ]; then count=1; else count=$(count); fi; \
	files=$$((count*2)); \
	ls -t migrations | head -n $$files | xargs -I {} rm migrations/{}

build:
	go build ./...

lint:
	golangci-lint run ./...

fmt:
	go fmt ./...

imports:
	goimports -w .

test:
	go test ./...

tidy:
	go mod tidy

lint-ci:
	golangci-lint run ./...

check: fmt imports lint

check-all: fmt imports lint build test tidy
