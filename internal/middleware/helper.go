package middleware

import (
	"expense-application/internal/model"
	"github.com/gin-gonic/gin"
)

func getPermissions(c *gin.Context) []string {
	permissions, _ := c.Get("user")
	return permissions.(model.User).Roles[0].PermissionsUnmarshalled
}
