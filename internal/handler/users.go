package handler

import (
	"expense-application/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetCurrentUser
// @Security ApiKeyAuth
// @Tags Users
// @Description Method allow to get current user
// @ID users-current
// @Accept json
// @Produce json
// @Success 200 {object} UserResponse
// @Router /users/current [get]
func (h *Handler) GetCurrentUser(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"data": model.UserResponse{
			Id:    user.(model.User).Id,
			Name:  user.(model.User).Name,
			TgId:  uint64(user.(model.User).TgId),
			Email: user.(model.User).Email,
			Role:  user.(model.User).Roles[0],
		},
	})
}

// UpdateUser
// @Security ApiKeyAuth
// @Tags Users
// @Param id path int true "Users ID"
// @Param form body UserUpdateRequest true "User form"
// @Description Method allow to update users
// @ID users-update
// @Accept json
// @Produce json
// @Success 200 {object} StatusResponse
// @Router /users/{id} [put]
func (h *Handler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64) // user_id

	currentUser, _ := c.Get("user")
	var user model.User

	err = c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if currentUser.(model.User).Id != uint(id) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}

	if id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "ID is not a type of int",
		})
		return
	}

	err = h.services.User.Update(&user, user.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success updated!",
	})
}
