package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
)

func UsersRead(c *gin.Context) {
	if slices.Contains(getPermissions(c), "users_read") {
		c.Next()
		return
	}
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
}

func UsersUpdate(c *gin.Context) {
	if slices.Contains(getPermissions(c), "users_update") {
		c.Next()
		return
	}
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
}
