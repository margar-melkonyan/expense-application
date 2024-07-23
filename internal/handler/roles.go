package handler

import (
	"expense-application/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetRoles(c *gin.Context) {
	roles, err := h.services.Role.Roles()

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": roles,
	})
}

func (h *Handler) GetRole(c *gin.Context) {
	var role model.Role
	roleID, _ := strconv.ParseUint(c.Param("id"), 10, 0)

	if roleID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Role ID must be provided",
		})
		return
	}

	role, _ = h.services.Role.Role(uint(roleID))

	c.JSON(http.StatusOK, gin.H{
		"data": role,
	})
}

func (h *Handler) StoreRole(c *gin.Context) {
	var role model.Role
	err := c.BindJSON(&role)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
	}

	err = h.services.Role.StoreRole(&role)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Couldn't created role",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success created!",
	})
}

func (h *Handler) UpdateRole(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 0)

	if id == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Id isn't a number",
		})
		return
	}

	var role model.Role

	err := c.BindJSON(&role)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.services.Role.UpdateRole(&role, uint(id))

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Couldn't updated role",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success updated!",
	})
}

func (h *Handler) DeleteRole(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 0)

	if id == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Id isn't a number",
		})
	}

	err := h.services.Role.DeleteRole(uint(id))

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success deleted!",
	})
}

func (h *Handler) AssignRoleToUser(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 0)

	if id == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Id isn't a number",
		})
		return
	}

	var userRole model.UserRole
	err := c.ShouldBindJSON(&userRole)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	userRole.UserID = uint(id)

	err = h.services.AssignRole(userRole)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Couldn't assign role to user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success assigned role to user",
	})
}

func (h *Handler) GetPermissions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": h.services.Role.Permissions(),
	})
}
