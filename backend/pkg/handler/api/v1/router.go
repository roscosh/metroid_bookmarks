package v1

import (
	"github.com/gin-gonic/gin"
	"metroid_bookmarks/pkg/handler/api/base_api"
	"metroid_bookmarks/pkg/handler/api/v1/areas"
	"metroid_bookmarks/pkg/handler/api/v1/auth"
	"metroid_bookmarks/pkg/handler/api/v1/bookmarks"
	"metroid_bookmarks/pkg/handler/api/v1/rooms"
	"metroid_bookmarks/pkg/handler/api/v1/skills"
	"metroid_bookmarks/pkg/handler/api/v1/users"
	"metroid_bookmarks/pkg/service"
)

type router struct {
	*baseApi.Router
	service *service.Service
}

func NewRouter(
	baseAPIRouter *baseApi.Router,
	service *service.Service,
) baseApi.ApiRouter {
	return &router{
		Router:  baseAPIRouter,
		service: service,
	}
}

func (h *router) RegisterHandlers(router *gin.RouterGroup) {
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
	bookmarksRouter := bookmarks.NewRouter(h.Router, h.service.Bookmarks)
	bookmarksRouter.RegisterHandlers(bookmarksGroup)
}
