include .env.development

.PHONY: up down dev build
init:
	go install github.com/air-verse/air
	go install github.com/golang/mock/mockgen@latest
	go install github.com/rakyll/gotest@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install entgo.io/ent/cmd/ent@latest

dev: # サーバー起動
	air

build:
	go build .

.PHONY: down migrations migrate migrate-down seed

up:
	docker compose up -d
down:
	docker-compose down
	rm -rf database/data

migrations: # スキーマ変更による差分からマイグレーションファイルを生成
	cd database; go run ./diff.go

migrate: # golang-migrate
	migrate -database "$(POSTGRESQL_URL)" -path database/migrations up

migrate-down: # golang-migrate
	migrate -database "$(POSTGRESQL_URL)" -path database/migrations down -all

seed: # テストデータを投入する
	docker exec -it template-go-api \
	psql -f ./docker-entrypoint-initdb.d/seed/inquiries.sql -U postgres -d postgres -h 0.0.0.0 -p 5432 \
	psql -f ./docker-entrypoint-initdb.d/seed/companies.sql -U postgres -d postgres -h 0.0.0.0 -p 5432 \
	psql -f ./docker-entrypoint-initdb.d/seed/users.sql -U postgres -d postgres -h 0.0.0.0 -p 5432

# conn-db: # ローカル環境からdbコンテナに接続
# 	PGPASSWORD=postgres psql -h localhost -p 4432 -U postgres -d postgres