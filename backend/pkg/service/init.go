package service

import (
	"metroid_bookmarks/misc"
	"metroid_bookmarks/pkg/repository/redis"
	"metroid_bookmarks/pkg/repository/sql"
)

var logger = misc.GetLogger()

type Service struct {
	Middleware    *MiddlewareService
	Authorization *AuthService
	Users         *UsersService
	Areas         *AreasService
	Rooms         *RoomsService
	Skills        *SkillsService
	Bookmarks     *BookmarksService
	Photos        *PhotosService
}

func NewService(sql *sql.SQL, redis *redis.Redis) *Service {
	return &Service{
		Middleware:    newMiddlewareService(sql.Users, redis.Session),
		Authorization: newAuthService(sql.Users),
		Users:         newUsersService(sql.Users),
		Areas:         newAreasService(sql.Areas),
		Rooms:         newRoomsService(sql.Rooms),
		Skills:        newSkillsService(sql.Skills),
		Bookmarks:     newBookmarksService(sql.Bookmarks, sql.Photos),
		Photos:        newPhotosService(sql.Photos),
	}
}
