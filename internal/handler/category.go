package handler

import (
	"expense-application/internal/dto/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetCategories(c *gin.Context) {
	categories, err := h.services.Category.IndexCategories()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
	}

	c.JSON(http.StatusOK, map[string][]response.Category{
		"data": categories,
	})
}
