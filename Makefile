DB_DSN?="postgres://postgres:secret@localhost:5432/postgres?sslmode=disable"
MIGRATIONS_DIR?="./migrations"

.SILENT:

restart: down up

up: docker_up

down: docker_down

docker_up:
	docker compose build && docker compose up -d

docker_down:
	docker compose down

migrate_up:
	migrate -verbose -path ${MIGRATIONS_DIR} -database ${DB_DSN} up

migrate_down:
	migrate -verbose -path ${MIGRATIONS_DIR} -database ${DB_DSN} down

migrate_create:
	migrate create -dir ${MIGRATIONS_DIR} -ext sql -seq ${name}

migrate_version:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_DSN} version

test:
	PG_DSN=${DB_DSN} BASE_PATH=http://localhost:8000/api \
	go test -v -count=1 ./...
