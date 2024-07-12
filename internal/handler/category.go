package handler

import (
	"expense-application/internal/dto/request"
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
	var category model.Category
	types := []string{"expense", "income"}

	err := c.ShouldBindJSON(&category)

	if !slices.Contains(types, category.Type) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Type is not income or expense"})
		return
	}

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	_, err = h.services.Category.Store(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't save it"})
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "Success created!",
	})
}

func (h *Handler) UpdateCategory(c *gin.Context) {
	slug := c.Param("slug")
	var category request.Category

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
	slug := c.Param("slug")
	_, err := h.services.Category.Delete(slug)

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success updated!",
	})
}
