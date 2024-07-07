package handler

import (
	"expense-application/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		categories := api.Group("/categories")
		{
			categories.GET(
				"/",
				h.getCategories,
			)
			categories.POST(
				"/",
				h.storeCategory,
			)
		}
	}

	return router
}
