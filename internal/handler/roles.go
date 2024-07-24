package handler

import (
	"expense-application/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetRole
// @Security ApiKeyAuth[admin]
// @Tags Roles
// @Param id path int true "Role ID"
// @Description Method that return role by ID
// @ID get-role
// @Accept json
// @Produce json
// @Success 200 {object} RoleResponse
// @Router /roles/{id} [get]
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

// GetRoles
// @Security ApiKeyAuth[admin]
// @Tags Roles
// @Description Method that return list of roles
// @ID get-roles
// @Accept json
// @Produce json
// @Success 200 {object} RolesResponse
// @Router /roles [get]
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

// StoreRole
// @Security ApiKeyAuth[admin]
// @Tags Roles
// @Param form body RoleResponseRequest true "Role form"
// @Description Method that store role
// @ID store-roles
// @Accept json
// @Produce json
// @Success 200 {object} StatusResponse
// @Router /roles [post]
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

// UpdateRole
// @Security ApiKeyAuth[admin]
// @Tags Roles
// @Param id path int true "Role ID"
// @Param form body RoleResponseRequest true "Role form"
// @Description Method that allow to update role by ID
// @ID update-roles
// @Accept json
// @Produce json
// @Success 200 {object} StatusResponse
// @Failure 422 {object} ErrorResponse
// @Router /roles/{id} [put]
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

// DeleteRole
// @Security ApiKeyAuth[admin]
// @Tags Roles
// @Param id path int true "Role ID"
// @Description Method that allow to delete role by ID
// @ID delete-roles
// @Accept json
// @Produce json
// @Success 200 {object} StatusResponse
// @Failure 422 {object} ErrorResponse
// @Router /roles/{id} [delete]
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

// AssignRoleToUser
// @Security ApiKeyAuth[admin]
// @Tags Roles
// @Param id path int true "User ID"
// @Param form body AssignRoleToUserRequest true "Role form"
// @Description Method that return list of permissions
// @ID assign-role-to-users
// @Accept json
// @Produce json
// @Success 200 {object} StatusResponse
// @Router  /users/:id/assign-role  [put]
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

// GetPermissions
// @Security ApiKeyAuth[admin]
// @Tags Roles
// @Description Method that return list of permissions
// @ID get-permissions
// @Accept json
// @Produce json
// @Success 200 {object} PermissionsResponse
// @Router /roles/permissions [get]
func (h *Handler) GetPermissions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": h.services.Role.Permissions(),
	})
}
