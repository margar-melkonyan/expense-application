package handler

import (
	"expense-application/internal/helper"
	"expense-application/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
)

// GetCategory
// @Security ApiKeyAuth
// @Tags Categories
// @Param id path int true "CategoryRequest ID"
// @Description Method that return category by id
// @ID category-get
// @Accept json
// @Produce json
// @Success 200 {object} CategoryResponse
// @Router /categories/{id} [get]
func (h *Handler) GetCategory(c *gin.Context) {
	slug := c.Param("slug")
	category, err := h.services.Category.GetCategoryBySlug(slug)

	if err != nil || category.Name == "" && category.Type == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "CategoryRequest not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// CategoryList
// @Security ApiKeyAuth
// @Tags Categories
// @Description Method that return list of categories
// @ID categories-get
// @Accept json
// @Produce json
// @Success 200 {object} CategoriesResponse
// @Router /categories [get]
func (h *Handler) CategoryList(c *gin.Context) {
	categories, err := h.services.Category.IndexCategories()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categories,
	})
}

// StoreCategory
// @Security ApiKeyAuth[admin]
// @Tags Categories
// @Param form body CategoryRequest true "CategoryRequest form"
// @Description Method that store category
// @ID categories-store
// @Accept json
// @Produce json
// @Success 200 {object} StatusResponse
// @Router /categories [post]
func (h *Handler) StoreCategory(c *gin.Context) {
	var category model.Category
	types := []string{"expense", "income"}

	err := c.ShouldBindJSON(&category)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, helper.FormatValidationError(err))
		return
	}

	if !slices.Contains(types, category.Type) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Type is not income or expense"})
		return
	}

	_, err = h.services.Category.Store(category)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, helper.FormatValidationError(err))
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "Success created!",
	})
}

// UpdateCategory
// @Security ApiKeyAuth[admin]
// @Tags Categories
// @Param id path int true "CategoryRequest ID"
// @Param form body CategoryRequest true "CategoryRequest form"
// @Description Method that update category by ID
// @ID categories-update
// @Accept json
// @Produce json
// @Success 200 {object} StatusResponse
// @Router /categories/{id} [put]
func (h *Handler) UpdateCategory(c *gin.Context) {
	slug := c.Param("slug")
	var category model.Category

	err := c.ShouldBindJSON(&category)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, helper.FormatValidationError(err))
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

// DeleteCategory
// @Security ApiKeyAuth[admin]
// @Tags Categories
// @Param id path int true "CategoryRequest ID"
// @Description Method that delete category by ID
// @ID categories-delete
// @Accept json
// @Produce json
// @Success 200 {object} StatusResponse
// @Router /categories/{id} [delete]
func (h *Handler) DeleteCategory(c *gin.Context) {
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
