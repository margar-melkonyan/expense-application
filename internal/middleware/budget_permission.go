package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
)

func BudgetsCreate(c *gin.Context) {
	if slices.Contains(getPermissions(c), "budgets_create") {
		c.Next()
		return
	}
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
}

func BudgetsRead(c *gin.Context) {
	if slices.Contains(getPermissions(c), "budgets_read") {
		c.Next()
		return
	}
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
}

func BudgetsUpdate(c *gin.Context) {
	if slices.Contains(getPermissions(c), "budgets_update") {
		c.Next()
		return
	}
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
}

func BudgetsDelete(c *gin.Context) {
	if slices.Contains(getPermissions(c), "budgets_delete") {
		c.Next()
		return
	}
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
}
