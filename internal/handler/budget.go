package handler

import (
	"expense-application/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
	"strconv"
)

func (h *Handler) GetUserBudget(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "ID is not a type of int",
		})
		return
	}

	budgets, err := h.services.Budget.GetUserBudgets(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": budgets})
}

func (h *Handler) StoreBudget(c *gin.Context) {
	var budget model.Budget
	var category model.Category
	types := []string{"expense", "income"}

	err := c.ShouldBindJSON(&budget)

	category, _ = h.services.Category.GetCategoryBySlug(budget.CategorySlug)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if !slices.Contains(types, budget.Type) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Type is not income or expense"})
		return
	}

	err = h.services.Budget.Store(budget, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success created!",
	})
}

func (h *Handler) DeleteBudget(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "ID is not a type of int",
		})
		return
	}

	_, err = h.services.Budget.Delete(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success deleted!",
	})
}
