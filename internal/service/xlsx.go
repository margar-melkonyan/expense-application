package service

import (
	"bytes"
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

func (s XLSXService) GenDayReport(typeBudget string, userId int) *bytes.Buffer {
	file := excelize.NewFile()

	headers := []string{"ID", "Имя", "Возраст"}
	for i, header := range headers {
		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), 1), header)
	}

	data := [][]interface{}{
		{1, "John", 30},
		{2, "Alex", 20},
		{3, "Bob", 40},
	}

	for i, row := range data {
		dataRow := i + 2
		for j, col := range row {
			file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+j)), dataRow), col)
		}
	}

	buffer, _ := file.WriteToBuffer()

	return buffer
}
func (s XLSXService) GenWeekReport(typeBudget string, userId int) *bytes.Buffer {
	file := excelize.NewFile()

	headers := []string{"ID", "Имя", "Возраст"}
	for i, header := range headers {
		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), 1), header)
	}

	data := [][]interface{}{
		{1, "John", 30},
		{2, "Alex", 20},
		{3, "Bob", 40},
	}

	for i, row := range data {
		dataRow := i + 2
		for j, col := range row {
			file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+j)), dataRow), col)
		}
	}

	buffer, _ := file.WriteToBuffer()

	return buffer
}

func (s XLSXService) GenMonthReport(typeBudget string, userId int) *bytes.Buffer {
	file := excelize.NewFile()

	headers := []string{"ID", "Имя", "Возраст"}
	for i, header := range headers {
		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), 1), header)
	}

	data := [][]interface{}{
		{1, "John", 30},
		{2, "Alex", 20},
		{3, "Bob", 40},
	}

	for i, row := range data {
		dataRow := i + 2
		for j, col := range row {
			file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+j)), dataRow), col)
		}
	}

	buffer, _ := file.WriteToBuffer()

	return buffer
}
