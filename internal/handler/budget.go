package handler

import (
	"expense-application/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
	"strconv"
)

// GetBudget
// @Security ApiKeyAuth[admin]
// @Tags Budgets
// @Param id path int true "Budget ID"
// @Description Method that return budget by ID
// @ID get-budget
// @Accept json
// @Produce json
// @Success 200 {object} StatusResponse
// @Router /budgets/{id} [get]
func (h *Handler) GetBudget(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64) // budget_id
	if id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "ID is not a type of int",
		})
		return
	}

	budget, err := h.services.Budget.GetBudget(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, budget)
}

// GetUserBudgetList
// @Security ApiKeyAuth[admin]
// @Tags Budgets
// @Description Method that return list of budgets
// @ID get-budgets
// @Accept json
// @Produce json
// @Success 200 {object} StatusResponse
// @Router /budgets [get]
func (h *Handler) GetUserBudgetList(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64) // user_id

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

// StoreBudget
// @Security ApiKeyAuth[admin]
// @Tags Budgets
// @Param form body RoleResponseRequest true "Budget form"
// @Description Method that store role
// @ID store-budgets
// @Accept json
// @Produce json
// @Success 200 {object} StatusResponse
// @Router /budgets [post]
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

// UpdateBudget
// @Security ApiKeyAuth[admin]
// @Tags Budgets
// @Param id path int true "Budget ID"
// @Param form body RoleResponseRequest true "Budget form"
// @Description Method that allow to update budget by ID
// @ID update-budgets
// @Accept json
// @Produce json
// @Success 200 {object} StatusResponse
// @Failure 422 {object} ErrorResponse
// @Router /budgets/{id} [put]
func (h *Handler) UpdateBudget(c *gin.Context) {
	var budget model.Budget
	_, err := strconv.ParseUint(c.Param("id"), 10, 64) // budget_id

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "ID is not a type of int",
		})
		return
	}

	err = c.BindJSON(&budget)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	_, err = h.services.Budget.Update(budget)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success updated!"})
}

// DeleteBudget
// @Security ApiKeyAuth[admin]
// @Tags Budgets
// @Param id path int true "Budget ID"
// @Description Method that allow to delete budget by ID
// @ID delete-budgets
// @Accept json
// @Produce json
// @Success 200 {object} StatusResponse
// @Router /budgets/{id} [delete]
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
