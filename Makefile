include .env
export

DB_URL=postgres//$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable


migrate_up:
	migrate -path migrations -database $(DB_URL) up

migrate_down:
	migrate -path migrations -database $(DB_URL) down

create_migration:
	migrate create -ext sql -dir migrations $(name)

docker_build:
	docker build --platform linux/arm64 -t agro_storage -f ./build/Dockerfile .

compose_up:
	docker compose -f ./build/docker-compose.yaml up -d

compose_down:
	docker compose -f ./build/docker-compose.yaml down

compose_up_build:
	docker compose -f ./build/docker-compose.yaml up --build

