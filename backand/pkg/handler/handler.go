package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	_ "metroid_bookmarks/docs"
	"metroid_bookmarks/misc"
	"metroid_bookmarks/pkg/service"
)

var logger = misc.GetLogger()

type Handler struct {
	services *service.Service
	config   *misc.Config
}

func NewHandler(services *service.Service, config *misc.Config) *Handler {
	return &Handler{services: services, config: config}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1", h.getSession)

	auth := api.Group("/auth")
	{
		auth.GET("/me", h.me)
		auth.POST("/sign_up", h.signUp) //ТОЛЬКО ДЛЯ ТЕСТОВ, на продакшене использовать /users/create
		auth.POST("/login", h.login)
		auth.POST("/logout", h.authRequired, h.logout)
	}
	users := api.Group("/users", h.authRequired)
	{
		users.GET("/get_all", h.getAllUsers)
		users.POST("/create", h.adminRequired, h.createUser)
		users.DELETE("/delete/:id", h.adminRequired, h.deleteUser)
		users.PUT("/edit/:id", h.adminRequired, h.editUser)
		users.PUT("/change_password/:id", h.adminRequired, h.changePassword)
	}
	return router
}
