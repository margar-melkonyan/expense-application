package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
)

func CategoriesCreate(c *gin.Context) {
	if slices.Contains(getPermissions(c), "categories_create") {
		c.Next()
		return
	}
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
}

func CategoriesRead(c *gin.Context) {
	if slices.Contains(getPermissions(c), "categories_read") {
		c.Next()
		return
	}
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
}

func CategoriesUpdate(c *gin.Context) {
	if slices.Contains(getPermissions(c), "categories_update") {
		c.Next()
		return
	}
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
}

func CategoriesDelete(c *gin.Context) {
	if slices.Contains(getPermissions(c), "categories_delete") {
		c.Next()
		return
	}
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
}
