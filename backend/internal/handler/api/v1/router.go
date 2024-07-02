package v1

import (
	"metroid_bookmarks/internal/handler/api/middleware"
	"metroid_bookmarks/internal/handler/api/v1/areas"
	"metroid_bookmarks/internal/handler/api/v1/auth"
	"metroid_bookmarks/internal/handler/api/v1/bookmarks"
	"metroid_bookmarks/internal/handler/api/v1/photos"
	"metroid_bookmarks/internal/handler/api/v1/rooms"
	"metroid_bookmarks/internal/handler/api/v1/skills"
	"metroid_bookmarks/internal/handler/api/v1/users"
	"metroid_bookmarks/internal/service"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*middleware.Router
	service *service.Service
}

func NewRouter(
	mwRouter *middleware.Router,
	service *service.Service,
) *Router {
	return &Router{
		Router:  mwRouter,
		service: service,
	}
}

func (r *Router) RegisterHandlers(router *gin.RouterGroup) {
	authGroup := router.Group("/auth")
	authRouter := auth.NewRouter(r.Router, r.service.Authorization)
	authRouter.RegisterHandlers(authGroup)

	usersGroup := router.Group("/users", r.Middleware.AdminRequired)
	usersRouter := users.NewRouter(r.Router, r.service.Users)
	usersRouter.RegisterHandlers(usersGroup)

	areasGroup := router.Group("/areas", r.Middleware.AuthRequired)
	areasRouter := areas.NewRouter(r.Router, r.service.Areas)
	areasRouter.RegisterHandlers(areasGroup)

	roomsGroup := router.Group("/rooms", r.Middleware.AuthRequired)
	roomsRouter := rooms.NewRouter(r.Router, r.service.Rooms)
	roomsRouter.RegisterHandlers(roomsGroup)

	skillsGroup := router.Group("/skills", r.Middleware.AuthRequired)
	skillsRouter := skills.NewRouter(r.Router, r.service.Skills)
	skillsRouter.RegisterHandlers(skillsGroup)

	bookmarksGroup := router.Group("/bookmarks", r.Middleware.AuthRequired)
	bookmarksRouter := bookmarks.NewRouter(r.Router, r.service.Bookmarks, r.service.Photos)
	bookmarksRouter.RegisterHandlers(bookmarksGroup)

	photosGroup := router.Group("/photos", r.Middleware.AuthRequired)
	photosRouter := photos.NewRouter(r.Router, r.service.Photos, r.service.Bookmarks)
	photosRouter.RegisterHandlers(photosGroup)
}
