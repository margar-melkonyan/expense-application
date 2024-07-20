package handler

import (
	"expense-application/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
)

func (h *Handler) GetCategory(c *gin.Context) {
	slug := c.Param("slug")
	category, err := h.services.Category.GetCategoryBySlug(slug)

	if err != nil || category.Name == "" && category.Type == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "Category not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

func (h *Handler) CategoryList(c *gin.Context) {
	categories, err := h.services.Category.IndexCategories()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categories,
	})
}

func (h *Handler) StoreCategory(c *gin.Context) {
	if c.GetHeader("Authorization") != "" {
		c.JSON(http.StatusForbidden, gin.H{})
		return
	}

	var category model.Category
	types := []string{"expense", "income"}

	err := c.ShouldBindJSON(&category)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if !slices.Contains(types, category.Type) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Type is not income or expense"})
		return
	}

	_, err = h.services.Category.Store(category)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "Success created!",
	})
}

func (h *Handler) UpdateCategory(c *gin.Context) {
	if c.GetHeader("Authorization") != "" {
		c.JSON(http.StatusForbidden, gin.H{})
		return
	}

	slug := c.Param("slug")
	var category model.Category

	err := c.ShouldBindJSON(&category)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	_, err = h.services.Category.Update(slug, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't update it"})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success updated!",
	})
}

func (h *Handler) DeleteCategory(c *gin.Context) {
	if c.GetHeader("Authorization") != "" {
		c.JSON(http.StatusForbidden, gin.H{})
		return
	}

	slug := c.Param("slug")
	_, err := h.services.Category.Delete(slug)

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success deleted!",
	})
}
