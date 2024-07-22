package handler

import (
	"expense-application/internal/model"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

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

func (h *Handler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64) // user_id

	if err != nil {
		slog.Error(err.Error())
	}

	if id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "ID is not a type of int",
		})
		return
	}

	var user model.User

	err = c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
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
