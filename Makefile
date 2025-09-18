include .env
export $(shell sed -n 's/^\([A-Za-z0-9_]\+\)=.*/\1/p' .env)

DB_URL=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable

start-app:
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