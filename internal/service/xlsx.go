package service

import (
	"bytes"
	"expense-application/internal/consts"
	"expense-application/internal/consts/xlsx"
	"expense-application/internal/model"
	"expense-application/internal/repository"
	"fmt"
	"github.com/xuri/excelize/v2"
)

type XLSXService struct {
	budgetRepository repository.Budget
}

func NewXLSXService(budgetRepository repository.Budget) *XLSXService {
	return &XLSXService{
		budgetRepository: budgetRepository,
	}
}

func addTotalToXLSX(file *excelize.File, total float64, rowNumber int, stylesID []int) {
	file.MergeCell(
		"Sheet1",
		fmt.Sprintf("%s%d", string(rune(xlsx.D)), rowNumber),
		fmt.Sprintf("%s%d", string(rune(xlsx.E)), rowNumber),
	)
	file.SetCellStyle(
		"Sheet1",
		fmt.Sprintf("%s%d", string(rune(xlsx.D)), rowNumber),
		fmt.Sprintf("%s%d", string(rune(xlsx.E)), rowNumber),
		stylesID[0],
	)
	file.SetCellValue(
		"Sheet1",
		fmt.Sprintf("%s%d", string(rune(xlsx.D)), rowNumber),
		"Category total :",
	)

	file.MergeCell(
		"Sheet1",
		fmt.Sprintf("%s%d", string(rune(xlsx.F)), rowNumber),
		fmt.Sprintf("%s%d", string(rune(xlsx.G)), rowNumber),
	)
	file.SetCellStyle(
		"Sheet1",
		fmt.Sprintf("%s%d", string(rune(xlsx.F)), rowNumber),
		fmt.Sprintf("%s%d", string(rune(xlsx.G)), rowNumber),
		stylesID[1],
	)
	file.SetCellValue(
		"Sheet1",
		fmt.Sprintf("%s%d", string(rune(xlsx.F)), rowNumber),
		fmt.Sprintf("%.2f", total),
	)

	file.SetCellStyle(
		"Sheet1",
		fmt.Sprintf("%s%d", string(rune(xlsx.H)), rowNumber),
		fmt.Sprintf("%s%d", string(rune(xlsx.H)), rowNumber),
		stylesID[0],
	)
	file.SetCellValue(
		"Sheet1",
		fmt.Sprintf("%s%d", string(rune(xlsx.H)), rowNumber),
		"руб.",
	)
}

func genXLSX(header string, budgetCategories []model.Category) *bytes.Buffer {
	file := excelize.NewFile()

	textStyle, _ := file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})

	dateFormat, _ := file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		NumFmt: 15,
	})

	numberFormat, _ := file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		NumFmt: 2,
	})

	file.MergeCell("Sheet1", "A1", "H1")
	file.SetCellStyle("Sheet1", "A1", "H1", textStyle)
	file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(xlsx.A)), 1), header)

	rowNumber := 0
	totalSum := 0.0

	for _, category := range budgetCategories {
		if len(category.Budgets) != 0 {
			rowNumber += 3

			file.MergeCell("Sheet1", fmt.Sprintf(
				"%s%d",
				string(rune(xlsx.A)), rowNumber),
				fmt.Sprintf("%s%d", string(rune(xlsx.H)), rowNumber),
			)
			file.SetCellStyle("Sheet1", fmt.Sprintf(
				"%s%d",
				string(rune(xlsx.A)), rowNumber),
				fmt.Sprintf("%s%d", string(rune(xlsx.H)), rowNumber),
				textStyle,
			)
			file.SetCellValue("Sheet1", fmt.Sprintf(
				"%s%d",
				string(rune(xlsx.A)), rowNumber),
				category.Name,
			)

			categorySum := 0.0
			for j, budget := range category.Budgets {
				rowNumber += 1

				file.SetCellValue(
					"Sheet1",
					fmt.Sprintf("%s%d", string(rune(xlsx.A)), rowNumber),
					j+1,
				)

				//Title Cell
				file.SetCellValue(
					"Sheet1",
					fmt.Sprintf("%s%d", string(rune(xlsx.B)), rowNumber),
					budget.Title,
				)
				file.MergeCell(
					"Sheet1",
					fmt.Sprintf("%s%d", string(rune(xlsx.B)), rowNumber),
					fmt.Sprintf("%s%d", string(rune(xlsx.E)), rowNumber),
				)

				////Amount Cell
				file.MergeCell(
					"Sheet1",
					fmt.Sprintf("%s%d", string(rune(xlsx.F)), rowNumber),
					fmt.Sprintf("%s%d", string(rune(xlsx.G)), rowNumber),
				)
				file.SetCellStyle(
					"Sheet1",
					fmt.Sprintf("%s%d", string(rune(xlsx.F)), rowNumber),
					fmt.Sprintf("%s%d", string(rune(xlsx.G)), rowNumber),
					numberFormat,
				)
				file.SetCellValue(
					"Sheet1",
					fmt.Sprintf("%s%d", string(rune(xlsx.F)), rowNumber),
					fmt.Sprintf("%.2f", budget.Amount/100),
				)

				// CreatedAt Cell
				file.SetCellValue(
					"Sheet1",
					fmt.Sprintf("%s%d", string(rune(xlsx.H)), rowNumber),
					budget.CreatedAt,
				)
				file.SetCellStyle(
					"Sheet1",
					fmt.Sprintf("%s%d", string(rune(xlsx.H)), rowNumber),
					fmt.Sprintf("%s%d", string(rune(xlsx.H)), rowNumber),
					dateFormat,
				)

				categorySum += budget.Amount / 100
			}
			rowNumber += 1
			totalSum += categorySum
			addTotalToXLSX(file, categorySum, rowNumber, []int{textStyle, numberFormat})
		}
	}

	rowNumber += 4
	addTotalToXLSX(file, totalSum, rowNumber, []int{textStyle, numberFormat})
	buffer, _ := file.WriteToBuffer()

	return buffer
}

func (s XLSXService) GenDayReport(typeBudget string, userId uint) *bytes.Buffer {
	budgets, _ := s.budgetRepository.GetBudgetByCategoryAndPeriod(typeBudget, userId, consts.Day)
	return genXLSX(fmt.Sprintf("Current %s / %s", consts.Day, typeBudget), budgets)
}
func (s XLSXService) GenWeekReport(typeBudget string, userId uint) *bytes.Buffer {
	budgets, _ := s.budgetRepository.GetBudgetByCategoryAndPeriod(typeBudget, userId, consts.Week)
	return genXLSX(fmt.Sprintf("Current %s / %s", consts.Week, typeBudget), budgets)
}

func (s XLSXService) GenMonthReport(typeBudget string, userId uint) *bytes.Buffer {
	budgets, _ := s.budgetRepository.GetBudgetByCategoryAndPeriod(typeBudget, userId, consts.Month)
	return genXLSX(fmt.Sprintf("Current %s / %s", consts.Month, typeBudget), budgets)
}
