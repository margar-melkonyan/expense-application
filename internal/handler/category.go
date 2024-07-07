package handler

import (
	"errors"
	"expense-application/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
)

func (h *Handler) getCategories(ctx *gin.Context) {

}

func (h *Handler) storeCategory(ctx *gin.Context) {
	var category model.Category

	err := ctx.Bind(&category)
	if err != nil {
		slog.Error(err.Error())
	}

	id, err := h.services.Category.Store(category)

	if err != nil && errors.Is(err, gorm.ErrDuplicatedKey) {
		fmt.Println("test")
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
