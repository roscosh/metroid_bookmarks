include .env

export DOCKER_BUILDKIT=0
export COMPOSE_ENV_FILE=.env

ifeq ($(PRODUCTION), true)
	FILE=./conf/docker/prod/docker-compose.yml
	PROJECT=metroid_bookmarks_prod
else
	FILE=./conf/docker/dev/docker-compose.yml
	PROJECT=metroid_bookmarks_dev
endif

# build containers
build:
	docker compose -f $(FILE) --project-directory . -p $(PROJECT) build app

# Run containers
start:
	docker compose -f $(FILE) --project-directory . -p $(PROJECT) up app

# Run containers in the background
silence_start:
	docker compose -f $(FILE) --project-directory . -p $(PROJECT) up -d app

# show and follow logs of container
logs:
	docker compose -p $(PROJECT) logs app -f

# stop container
stop:
	docker compose -p $(PROJECT) stop

# completely remove application with container, image and volumes
clean:
	docker compose -p $(PROJECT) down --rmi "local" -v

# run pending postgreSQL migrations
migrate:
	docker compose -f $(FILE) --project-directory . -p $(PROJECT) up dbmate

# run linter(only for develop mode)
lint:
	docker compose -f ./conf/docker/dev/docker-compose.yml --project-directory . -p $(PROJECT) up linter

# run tests(only for develop mode)
tests:
	docker compose -f ./conf/docker/dev/docker-compose.yml --project-directory . -p $(PROJECT) up test

mockgen:
	docker compose -f ./conf/docker/dev/docker-compose.yml --project-directory . -p $(PROJECT) run --rm linter bash ./conf/mockgen.sh

.PHONY: cover
cover:
	go test -short -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out