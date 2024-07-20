package handler

import (
	"expense-application/internal/model"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) SignUp(c *gin.Context) {
	var user model.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		slog.Error(err.Error())
	}

	body, err := h.services.SignUp(&user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, body)
}

func (h *Handler) SignIn(c *gin.Context) {
	var user model.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		slog.Error(err.Error())
	}

	body, err := h.services.SignIn(&user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, body)
}

func (h *Handler) RefreshToken(c *gin.Context) {
	h.services.RefreshToken(c)
}

func (h *Handler) Logout(c *gin.Context) {
	c.Set("user", nil)
	c.JSON(http.StatusOK, gin.H{
		"message": "You logged out!",
	})
}
