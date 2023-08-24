create:
	migrate create -ext sql -dir migrations -seq $(name)

up:
	migrate -database ${POSTGRESQL_URL} -path migrations up

allDown:
	migrate -database ${POSTGRESQL_URL} -path migrations down

build:
	docker compose build

start:
	docker compose up

db:
	docker compose -f docker-compose.postgres.yml up -d