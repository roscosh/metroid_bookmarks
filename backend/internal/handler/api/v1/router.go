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

func (h *Router) RegisterHandlers(router *gin.RouterGroup) {
	authGroup := router.Group("/auth")
	authRouter := auth.NewRouter(h.Router, h.service.Authorization)
	authRouter.RegisterHandlers(authGroup)

	usersGroup := router.Group("/users", h.Middleware.AdminRequired)
	usersRouter := users.NewRouter(h.Router, h.service.Users)
	usersRouter.RegisterHandlers(usersGroup)

	areasGroup := router.Group("/areas", h.Middleware.AuthRequired)
	areasRouter := areas.NewRouter(h.Router, h.service.Areas)
	areasRouter.RegisterHandlers(areasGroup)

	roomsGroup := router.Group("/rooms", h.Middleware.AuthRequired)
	roomsRouter := rooms.NewRouter(h.Router, h.service.Rooms)
	roomsRouter.RegisterHandlers(roomsGroup)

	skillsGroup := router.Group("/skills", h.Middleware.AuthRequired)
	skillsRouter := skills.NewRouter(h.Router, h.service.Skills)
	skillsRouter.RegisterHandlers(skillsGroup)

	bookmarksGroup := router.Group("/bookmarks", h.Middleware.AuthRequired)
	bookmarksRouter := bookmarks.NewRouter(h.Router, h.service.Bookmarks, h.service.Photos)
	bookmarksRouter.RegisterHandlers(bookmarksGroup)

	photosGroup := router.Group("/photos", h.Middleware.AuthRequired)
	photosRouter := photos.NewRouter(h.Router, h.service.Photos, h.service.Bookmarks)
	photosRouter.RegisterHandlers(photosGroup)
}
