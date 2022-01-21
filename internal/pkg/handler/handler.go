package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ziyadovea/todo-app/internal/pkg/service"
)

type Handler struct {
	Services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		Services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.GET("/", h.getAllLists)
			lists.POST("/", h.createNewList)
			lists.DELETE("/", h.deleteAllLists)
			lists.GET("/:list_id", h.getListByID)
			lists.PUT("/:list_id", h.updateListByID)
			lists.DELETE("/:list_id", h.deleteListByID)

			items := api.Group(":id/items")
			{
				items.GET("/", h.getAllItems)
				items.POST("/", h.createNewItem)
				items.DELETE("/", h.deleteAllItems)
				items.GET("/:item_id", h.getItemByID)
				items.PUT("/:item_id", h.updateItemByID)
				items.DELETE("/:item_id", h.deleteItemByID)
			}
		}
	}

	return router
}
