package handler

import (
	"expense-application/internal/helper"
	"expense-application/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SignUp
// @Summary SignUp
// @Tags Auth
// @Description Method allow to create account
// @ID auth-create-account
// @Accept json
// @Produce json
// @Param form body SignUpRequest true "account info"
// @Success 200 {object} AuthResponse
// @Router /auth/sign-up [post]
func (h *Handler) SignUp(c *gin.Context) {
	var user model.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, helper.FormatValidationError(err))
		return
	}

	body, err := h.services.SignUp(&user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, body)
}

// SignIn
// @Summary SignIn
// @Tags Auth
// @Description Method allow to enter into account
// @ID auth-enter-account
// @Accept json
// @Produce json
// @Param input body SignInRequest true "account info"
// @Success 200 {object} AuthResponse
// @Router /auth/sign-in [post]
func (h *Handler) SignIn(c *gin.Context) {
	var user model.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, helper.FormatValidationError(err))
		return
	}

	body, err := h.services.SignIn(&user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, helper.FormatValidationError(err))
		return
	}

	c.JSON(http.StatusOK, body)
}

// RefreshToken
// @Summary RefreshToken
// @Security ApiKeyAuth
// @Tags Auth
// @Description Method that return new pair of access and refresh token
// @ID auth-refresh-token
// @Accept json
// @Produce json
// @Success 200 {object} AuthResponse
// @Router /auth/refresh-token [post]
func (h *Handler) RefreshToken(c *gin.Context) {
	h.services.RefreshToken(c)
}

// Logout
// @Summary Logout
// @Security ApiKeyAuth
// @Tags Auth
// @Description Method that logout user on server side
// @ID auth-logout
// @Accept json
// @Produce json
// @Success 200 {object} StatusResponse
// @Router /auth/logout [post]
func (h *Handler) Logout(c *gin.Context) {
	c.Set("user", nil)
	c.JSON(http.StatusOK, gin.H{
		"message": "You logged out!",
	})
}
