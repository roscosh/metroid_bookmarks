# Include environment variables from the .env file
include .env

# Disable Docker BuildKit (set to 0)
# This allows images to be loaded from the cache without checking metadata from Docker Hub
export DOCKER_BUILDKIT=0
# Specify the .env file for docker-compose to use
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

# Run containers in the background (detached mode)
silence_start:
	docker compose -f $(FILE) --project-directory . -p $(PROJECT) up -d app

# show and follow logs of the running app container
logs:
	docker compose -p $(PROJECT) logs app -f

# Stop running app containers
stop:
	docker compose -p $(PROJECT) stop

# completely remove application with container, image and volumes
clean:
	docker compose -p $(PROJECT) down --rmi "local" -v

# run any pending PostgreSQL migrations using the dbmate service
migrate:
	docker compose -f $(FILE) --project-directory . -p $(PROJECT) up dbmate

# run the linter(only for 'develop' mode)
lint:
	docker compose -f ./conf/docker/dev/docker-compose.yml --project-directory . -p $(PROJECT) up linter

# run Go tests with coverage(only for 'develop' mode)
.PHONY: test
test:
	docker compose -f ./conf/docker/dev/docker-compose.yml --project-directory . -p $(PROJECT) up test

# Generate mocks using mockgen(only for 'develop' mode)
mockgen:
	docker compose -f ./conf/docker/dev/docker-compose.yml --project-directory . -p $(PROJECT) run --rm linter bash ./conf/mockgen.sh