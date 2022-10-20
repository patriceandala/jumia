
migrate:
	docker run --network host migrate/migrate -path migrations -database postgres://root:secret@localhost:5433/jumia?sslmode=disable up

up:
	docker-compose up

down:
	docker-compose down
	docker run -v migrations:/migrations --network host migrate/migrate  -path=/migrations/ -database postgres://root:secret@localhost:5432/database down 2

createDb:
	docker exec  -it postgres15 createdb --username=root --owner=root jumia

.PHONY: server
server:
	$(MAKE) up
	$(MAKE) createDb
	$(MAKE) migrate


