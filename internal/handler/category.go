package handler

import (
	"expense-application/internal/model"
	"github.com/gin-gonic/gin"
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

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
