#!/bin/bash

#sql repository

mockgen \
	-source=./internal/repository/sql/areas/areas.go \
	-destination=./internal/repository/sql/areas/mocks/areas.go

mockgen \
	-source=./internal/repository/sql/bookmarks/bookmarks.go \
	-destination=./internal/repository/sql/bookmarks/mocks/bookmarks.go

mockgen \
	-source=./internal/repository/sql/photos/photos.go \
	-destination=./internal/repository/sql/photos/mocks/photos.go

mockgen \
	-source=./internal/repository/sql/rooms/rooms.go \
	-destination=./internal/repository/sql/rooms/mocks/rooms.go

mockgen \
	-source=./internal/repository/sql/skills/skills.go \
	-destination=./internal/repository/sql/skills/mocks/skills.go

mockgen \
	-source=./internal/repository/sql/users/users.go \
	-destination=./internal/repository/sql/users/mocks/users.go

#redis repository

mockgen \
	-source=./internal/repository/redis/session/session.go \
	-destination=./internal/repository/redis/session/mocks/session.go

#pgpool pkg
mockgen \
	-source=./pkg/pgpool/sql.go \
	-destination=./pkg/pgpool/mocks/sql.go