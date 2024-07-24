package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
)

func RolesCreate(c *gin.Context) {
	if slices.Contains(getPermissions(c), "roles_create") {
		c.Next()
		return
	}
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
}

func RolesRead(c *gin.Context) {
	if slices.Contains(getPermissions(c), "roles_read") {
		c.Next()
		return
	}
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
}

func RolesUpdate(c *gin.Context) {
	if slices.Contains(getPermissions(c), "roles_update") {
		c.Next()
		return
	}
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
}

func RolesDelete(c *gin.Context) {
	if slices.Contains(getPermissions(c), "roles_delete") {
		c.Next()
		return
	}
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
}
