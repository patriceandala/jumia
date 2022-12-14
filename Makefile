createMigrationFile:
	migrate create -ext sql -dir db/migration seq init_schema

migrateup:
	migrate -path db/migration/ -database postgresql://root:secret@localhost:5433/jumia?sslmode=disable -verbose up

migratedown:
	migrate -path db/migration/ -database postgresql://root:secret@localhost:5433/jumia?sslmode=disable -verbose down

up:
	docker-compose up

down:
	docker-compose down
	$(MAKE) migratedown

sqlc:
	sqlc generate

.PHONY: createDb
createDb:
	docker exec  -it postgres15 createdb --username=root --owner=root jumia

.PHONY: server
server:
	$(MAKE) up
	$(MAKE) createDb
	$(MAKE) migrateup


