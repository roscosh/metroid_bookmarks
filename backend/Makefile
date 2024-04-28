include .env

ifeq ($(PRODUCTION), true)
	FILE=./etc/docker/prod/docker-compose.yml
	PROJECT=metroid_bookmarks_prod
else
	FILE=./etc/docker/dev/docker-compose.yml
	PROJECT=metroid_bookmarks_dev
endif

build:
	docker compose -f $(FILE) --env-file .env --project-directory . -p $(PROJECT) build

start:
	docker compose -f $(FILE) --env-file .env --project-directory . -p $(PROJECT) up

stop:
	docker compose -f $(FILE) --project-directory . -p $(PROJECT) stop

clean:
	docker compose -f $(FILE) --project-directory . -p $(PROJECT) down --rmi "all" -v


