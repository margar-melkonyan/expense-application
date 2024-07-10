package service

import (
	"expense-application/internal/model"
	"expense-application/internal/repository"
	"fmt"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/pdfcpu/pdfcpu/pkg/font"
	"log"
	"log/slog"
)

const day = "day"
const week = "week"
const month = "month"
const year = "year"

type PDFService struct {
	budgetRepository repository.Budget
}

func NewPdfService(budgetRepository repository.Budget) *PDFService {
	return &PDFService{
		budgetRepository: budgetRepository,
	}
}

func getMaroto(header string, budgetCategories []model.Category) core.Maroto {
	var contentsRow []core.Row
	budgetSum := 0.0

	font.UserFontDir = "assets/fonts"
	err := font.LoadUserFonts()
	if err != nil {
		slog.Error(err.Error())
	}

	cfg := config.NewBuilder().
		WithPageNumber().
		WithLeftMargin(15).
		WithRightMargin(15).
		WithTopMargin(10).
		WithBottomMargin(20).
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	m.AddRows(
		text.NewRow(20, header, props.Text{
			Style: fontstyle.Bold,
			Align: align.Center,
			Size:  18,
		}),
	)

	for _, budgetCategory := range budgetCategories {
		if len(budgetCategory.Budgets) != 0 {
			contentsRow = append(contentsRow, row.New(8).Add(
				text.NewCol(12, fmt.Sprintf("%s", budgetCategory.Name), props.Text{
					Top:   5,
					Style: fontstyle.Bold,
					Align: align.Center,
					Size:  12,
				}),
			).WithStyle(&props.Cell{BackgroundColor: getDarkModerateBlue()}))

			categorySum := 0.0

			for id, budget := range budgetCategory.Budgets {
				r := row.New(6).Add(
					text.NewCol(1, fmt.Sprintf("%d", id+1), props.Text{Size: 10, Align: align.Center}),
					text.NewCol(7, fmt.Sprintf("%s", budget.Title), props.Text{Size: 10, Align: align.Left}),
					text.NewCol(2, fmt.Sprintf("%.2f руб.", budget.Amount/100), props.Text{Size: 10, Align: align.Center}),
					text.NewCol(2, fmt.Sprintf("%v", budget.CreatedAt.Format("02.01.2006")), props.Text{Size: 10, Align: align.Center}),
				)

				if id%2 == 0 {
					r.WithStyle(&props.Cell{BackgroundColor: getBluishCyan()})
				} else {
					r.WithStyle(&props.Cell{BackgroundColor: getLightCyan()})
				}

				categorySum += budget.Amount / 100
				contentsRow = append(contentsRow, r)
			}

			budgetSum += categorySum

			contentsRow = append(contentsRow,
				row.New(5).Add(
					col.New(12),
				),
				row.New(2).Add(
					col.New(7),
					line.NewCol(5),
				),
				row.New(12).Add(
					col.New(7),
					text.NewCol(2, "Category total:", props.Text{
						Style: fontstyle.Bold,
						Size:  10,
						Align: align.Right,
					}),
					text.NewCol(3, fmt.Sprintf("%.2f руб.", categorySum), props.Text{
						Style: fontstyle.Bold,
						Size:  10,
						Align: align.Center,
					}),
				))
		}
	}

	contentsRow = append(contentsRow,
		row.New(10).Add(
			col.New(12),
		),
		row.New(12).Add(
			col.New(7),
			text.NewCol(2, "Total:", props.Text{
				Style: fontstyle.Bold,
				Size:  12,
				Align: align.Center,
			}),
			text.NewCol(3, fmt.Sprintf("%.2f руб.", budgetSum), props.Text{
				Style: fontstyle.BoldItalic,
				Size:  12,
				Align: align.Center,
			}),
		))

	m.AddRows(contentsRow...)

	return m
}

func getDarkModerateBlue() *props.Color {
	return &props.Color{
		Red:   59,
		Green: 109,
		Blue:  144,
	}
}

func getLightCyan() *props.Color {
	return &props.Color{
		Red:   239,
		Green: 244,
		Blue:  247,
	}
}

func getBluishCyan() *props.Color {
	return &props.Color{
		Red:   211,
		Green: 240,
		Blue:  246,
	}
}

func (s *PDFService) GenDayReport(typeBudget string, userId int) core.Document {
	budgets, _ := s.budgetRepository.GetBudgetByCategoryAndPeriod(typeBudget, userId, day)
	m := getMaroto(fmt.Sprintf("Current %s / %s", day, typeBudget), budgets)

	//file := excelize.NewFile()
	//
	//headers := []string{"ID", "Имя", "Возраст"}
	//for i, header := range headers {
	//	file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), 1), header)
	//}
	//
	//data := [][]interface{}{
	//	{1, "John", 30},
	//	{2, "Alex", 20},
	//	{3, "Bob", 40},
	//}
	//
	//for i, row := range data {
	//	dataRow := i + 2
	//	for j, col := range row {
	//		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+j)), dataRow), col)
	//	}
	//}
	//
	//buffer, _ := file.WriteToBuffer()

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err)
	}

	return document
}

func (s *PDFService) GenWeekReport(typeBudget string, userId int) core.Document {
	budgets, _ := s.budgetRepository.GetBudgetByCategoryAndPeriod(typeBudget, userId, week)
	m := getMaroto(fmt.Sprintf("Current %s / %s", week, typeBudget), budgets)
	document, err := m.Generate()

	if err != nil {
		log.Fatal(err)
	}

	return document
}

func (s *PDFService) GenMonthReport(typeBudget string, userId int) core.Document {
	budgets, _ := s.budgetRepository.GetBudgetByCategoryAndPeriod(typeBudget, userId, month)
	m := getMaroto(fmt.Sprintf("Current %s / %s", month, typeBudget), budgets)

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err)
	}

	return document
}
