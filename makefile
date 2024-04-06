db_start:
	docker-compose up -d
db_clean:
	docker-compose down
# migrate library to initialize db schema for migration
# up script to make a forward change to the schema, migrations will run updated migration files sequentially
# down script to revert changes
migrate_init:
	migrate create -ext sql -dir db/migration -seq  init-schema
migrate_up:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/kodboris?sslmode=disable" -verbose up
migrate_down:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/kodboris?sslmode=disable" -verbose down 1
migrate_fix:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/kodboris?sslmode=disable"  force 1
# drop the entire database schema
migrate_drop:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/bank?sslmode=disable" drop
.PHONY: db_start migrate_up test server
sqlc_init:
	sqlc init
sqlc_generate:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
