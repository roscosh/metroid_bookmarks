include .env

export DOCKER_BUILDKIT=0

ifeq ($(PRODUCTION), true)
	FILE=./conf/docker/prod/docker-compose.yml
	PROJECT=metroid_bookmarks_prod
else
	FILE=./conf/docker/dev/docker-compose.yml
	PROJECT=metroid_bookmarks_dev
endif

build:
	docker compose -f $(FILE) --env-file .env --project-directory . -p $(PROJECT) build

start:
	docker compose -f $(FILE) --env-file .env --project-directory . -p $(PROJECT) up

silence_start:
	docker compose -f $(FILE) --env-file .env --project-directory . -p $(PROJECT) up -d

logs:
	docker compose -p $(PROJECT) logs app -f

stop:
	docker compose -p $(PROJECT) stop

clean:
	docker compose -p $(PROJECT) down --rmi "local" -v

lint:
	golangci-lint run cmd/main.go