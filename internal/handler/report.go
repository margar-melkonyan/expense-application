package handler

import (
	"bytes"
	"expense-application/internal/model"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

// GeneratePDFReport
// @Security ApiKeyAuth
// @Tags Reports
// @Description Method for generation PDF report
// @ID reports-pdf
// @Accept json
// @Produce mpfd
// @Success 200 {object} StatusResponse
// @Router /reports/pdf [get]
func (h *Handler) GeneratePDFReport(c *gin.Context) {
	budgetType := c.Query("budget_type")

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

	var file []byte
	user, _ := c.Get("user")

	pdfService, ok := h.services.Reports["pdf"]
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	if period == "day" {
		file = pdfService.GenDayReport(budgetType, user.(model.User).Id)
	}
	if period == "week" {
		file = pdfService.GenWeekReport(budgetType, user.(model.User).Id)
	}
	if period == "month" {
		file = pdfService.GenMonthReport(budgetType, user.(model.User).Id)
	}

	contentLength := len(file)
	contentType := "application/pdf"
	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="report.pdf"`,
	}
	c.DataFromReader(http.StatusOK, int64(contentLength), contentType, bytes.NewReader(file), extraHeaders)
}

// GenerateXLSXReport
// @Security ApiKeyAuth
// @Tags Reports
// @Description Method for generation XLSX report
// @ID reports-xlsx
// @Accept json
// @Produce mpfd
// @Success 200 {object} StatusResponse
// @Router /reports/xlsx [get]
func (h *Handler) GenerateXLSXReport(c *gin.Context) {
	budgetType := c.Query("budget_type")

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

	var file []byte
	user, _ := c.Get("user")
	xlsxService, ok := h.services.Reports["xlsx"]
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	if period == "day" {
		file = xlsxService.GenDayReport(budgetType, user.(model.User).Id)
	}
	if period == "week" {
		file = xlsxService.GenWeekReport(budgetType, user.(model.User).Id)
	}
	if period == "month" {
		file = xlsxService.GenMonthReport(budgetType, user.(model.User).Id)
	}

	contentLength := len(file)
	contentType := "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="report.xlsx"`,
	}
	c.DataFromReader(http.StatusOK, int64(contentLength), contentType, bytes.NewBuffer(file), extraHeaders)
}
