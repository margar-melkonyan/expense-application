package handler

import (
	"bytes"
	"expense-application/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"net/http"
	"slices"
)

func (h *Handler) GeneratePDFReport(c *gin.Context) {
	budgetType := c.Query("budget_type")
	fmt.Println(budgetType)
	if !slices.Contains([]string{"income", "expense"}, budgetType) && budgetType != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "budget type is incorrect",
		})
		return
	}

	period := c.Query("period")
	if !slices.Contains([]string{"day", "week", "month"}, period) && period != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "budget type is incorrect",
		})
		return
	}

	var file core.Document
	user, _ := c.Get("user")

	if period == "day" {
		file = h.services.PDF.GenDayReport(budgetType, user.(model.User).Id)
	}
	if period == "week" {
		file = h.services.PDF.GenWeekReport(budgetType, user.(model.User).Id)
	}
	if period == "month" {
		file = h.services.PDF.GenMonthReport(budgetType, user.(model.User).Id)
	}

	contentLength := len(file.GetBytes())
	contentType := "application/pdf"
	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="report.pdf"`,
	}

	c.DataFromReader(http.StatusOK, int64(contentLength), contentType, bytes.NewReader(file.GetBytes()), extraHeaders)
}

func (h *Handler) GenerateXLSXReport(c *gin.Context) {
	budgetType := c.Query("budget_type")
	fmt.Println(budgetType)
	if !slices.Contains([]string{"income", "expense"}, budgetType) && budgetType != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "budget type is incorrect",
		})
		return
	}

	period := c.Query("period")
	if !slices.Contains([]string{"day", "week", "month"}, period) && period != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "budget type is incorrect",
		})
		return
	}

	var file *bytes.Buffer
	user, _ := c.Get("user")

	if period == "day" {
		file = h.services.XLSX.GenDayReport(budgetType, user.(model.User).Id)
	}
	if period == "week" {
		file = h.services.XLSX.GenWeekReport(budgetType, user.(model.User).Id)
	}
	if period == "month" {
		file = h.services.XLSX.GenMonthReport(budgetType, user.(model.User).Id)
	}

	contentLength := file.Len()
	contentType := "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="report.xlsx"`,
	}

	c.DataFromReader(http.StatusOK, int64(contentLength), contentType, file, extraHeaders)
}
