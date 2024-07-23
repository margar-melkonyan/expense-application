package handler

import (
	"expense-application/internal/middleware"
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
	router.Use(gin.Logger())
	router.MaxMultipartMemory = 10 << 20 // 10 MiB files allow

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.SignUp)
			auth.POST("/sign-in", h.SignIn)
			auth.POST("/refresh-token", h.RefreshToken)
			auth.POST("/logout", middleware.RequireAuth, h.Logout)
		}

		categories := api.Group("/categories", middleware.RequireAuth)
		{
			categories.GET(
				"",
				middleware.CategoriesRead,
				h.CategoryList,
			)

			categories.GET(
				":slug",
				middleware.CategoriesRead,
				h.GetCategory,
			)

			categories.POST(
				"",
				middleware.CategoriesCreate,
				h.StoreCategory,
			)

			categories.PUT(
				":slug",
				middleware.CategoriesUpdate,
				h.UpdateCategory,
			)

			categories.DELETE(
				":slug",
				middleware.CategoriesDelete,
				h.DeleteCategory,
			)
		}

		budgets := api.Group("/budgets", middleware.RequireAuth)
		{
			budgets.GET(
				":id/user",
				middleware.BudgetsRead,
				h.GetUserBudgetList,
			)

			budgets.GET(
				":id",
				middleware.BudgetsRead,
				h.GetBudget,
			)

			budgets.POST(
				"",
				middleware.BudgetsCreate,
				h.StoreBudget,
			)

			budgets.PUT(
				":id",
				middleware.BudgetsUpdate,
				h.UpdateBudget,
			)

			budgets.DELETE(
				":id",
				middleware.BudgetsDelete,
				h.DeleteBudget,
			)
		}

		users := api.Group("/users", middleware.RequireAuth)
		{
			users.GET(
				"current",
				middleware.UsersRead,
				h.GetCurrentUser,
			)

			users.PUT(
				":id",
				middleware.UsersUpdate,
				h.UpdateUser,
			)

			users.PUT(":id/assign-role", h.AssignRoleToUser)
		}

		reports := api.Group("/reports")
		{
			reports.GET(
				"pdf",
				h.GeneratePDFReport,
			)

			reports.GET(
				"xlsx",
				h.GenerateXLSXReport,
			)
		}

		roles := api.Group("/roles", middleware.RequireAuth)
		{
			roles.GET("",
				middleware.RolesRead,
				h.GetRoles,
			)

			roles.GET(":id",
				middleware.RolesRead,
				h.GetRole,
			)

			roles.POST("",
				middleware.RolesCreate,
				h.StoreRole,
			)

			roles.PUT(":id",
				middleware.RolesUpdate,
				h.UpdateRole,
			)

			roles.DELETE(":id",
				middleware.RolesDelete,
				h.DeleteRole,
			)

			permissions := roles.Group("/permissions", middleware.RequireAuth)
			{
				permissions.GET("", h.GetPermissions)
			}
		}
	}

	return router
}
