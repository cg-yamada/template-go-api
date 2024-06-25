include .env.development

.PHONY: init up down dev build
init:
	go install github.com/air-verse/air
	go install github.com/golang/mock/mockgen@latest
	go install github.com/rakyll/gotest@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install entgo.io/ent/cmd/ent@latest

up:
	docker compose up -d
down:
	docker-compose down
	rm -rf database/data

dev: # サーバー起動
	air

build:
	go build .