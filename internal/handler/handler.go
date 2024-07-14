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
	router.MaxMultipartMemory = 10 << 20 // 10 MiB files allow

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.SignUp)
			auth.POST("/sign-in", h.SignIn)
			auth.POST("/refresh-token", h.RefreshToken)
		}

		categories := api.Group("/categories")
		{
			categories.GET(
				"",
				h.CategoryList,
			)

			categories.GET(
				":slug",
				h.GetCategory,
			)

			categories.POST(
				"",
				h.StoreCategory,
			)

			categories.PUT(
				":slug",
				h.UpdateCategory,
			)

			categories.DELETE(
				":slug",
				h.DeleteCategory,
			)
		}

		budgets := api.Group("/budgets")
		{
			budgets.GET(
				":id/user",
				h.GetUserBudgetList,
			)

			budgets.GET(
				":id",
				h.GetBudget,
			)

			budgets.POST(
				"",
				h.StoreBudget,
			)

			budgets.PUT(
				":id",
				h.UpdateBudget,
			)

			budgets.DELETE(
				":id",
				h.DeleteBudget,
			)
		}
	}

	return router
}
