createMigrationFile:
	migrate create -ext sql -dir db/migration seq init_schema

migrate:
	migrate -path db/migration/ -database postgresql://root:secret@localhost:5433/jumiaDB?sslmode=disable -verbose up

up:
	docker-compose up

down:
	docker-compose down
	docker run -v migrations:/migrations --network host migrate/migrate  -path=/migrations/ -database postgres://root:secret@localhost:5433/jumiaDB down 2

sqlc:
	sqlc generate

.PHONY: createDb
createDb:
	docker exec  -it postgres15 createdb --username=root --owner=root jumiaDB

.PHONY: server
server:
	$(MAKE) up
	$(MAKE) createDb
	$(MAKE) migrate


