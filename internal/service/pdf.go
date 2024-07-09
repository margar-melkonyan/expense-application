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
	"log"
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
	cfg := config.NewBuilder().
		WithPageNumber().
		WithLeftMargin(25).
		WithRightMargin(5).
		WithTopMargin(20).
		WithBottomMargin(20).
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	m.AddRows(
		text.NewRow(20, header, props.Text{
			Top:   3,
			Style: fontstyle.Bold,
			Align: align.Center,
			Size:  18,
		}),
	)

	m.AddRow(20, line.NewCol(12))

	var contentsRow []core.Row

	for _, budgetCategory := range budgetCategories {
		if len(budgetCategory.Budgets) != 0 {
			m.AddRows(
				row.New(5).Add(
					text.NewCol(10, budgetCategory.Name, props.Text{
						Top:   3,
						Style: fontstyle.Bold,
						Align: align.Center,
						Size:  18,
					}),
				),
			)

			for _, budget := range budgetCategory.Budgets {
				r := row.New(4).Add(
					col.New(3),
					text.NewCol(4, fmt.Sprintf("%s", budget.Title), props.Text{Size: 8, Align: align.Center}),
					text.NewCol(2, fmt.Sprintf("%.2f", budget.Amount/100), props.Text{Size: 8, Align: align.Center}),
					text.NewCol(3, fmt.Sprintf("%v", budget.CreatedAt), props.Text{Size: 8, Align: align.Center}),
				)

				contentsRow = append(contentsRow, r)
			}
		}
	}

	m.AddRows(contentsRow...)

	return m
}

func getDarkGrayColor() *props.Color {
	return &props.Color{
		Red:   55,
		Green: 55,
		Blue:  55,
	}
}

func getBlueColor() *props.Color {
	return &props.Color{
		Red:   10,
		Green: 10,
		Blue:  150,
	}
}

func getRedColor() *props.Color {
	return &props.Color{
		Red:   150,
		Green: 10,
		Blue:  10,
	}
}

func (s *PDFService) GenDayReport(typeBudget string, userId int) core.Document {
	budgets, _ := s.budgetRepository.GetBudgetByCategoryAndPeriod(typeBudget, userId, day)
	m := getMaroto(fmt.Sprintf("Current %s / %s", day, typeBudget), budgets)

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

func (s *PDFService) GenYearReport(typeBudget string, userId int) core.Document {
	budgets, _ := s.budgetRepository.GetBudgetByCategoryAndPeriod(typeBudget, userId, year)
	m := getMaroto(fmt.Sprintf("Current %s / %s", year, typeBudget), budgets)
	document, err := m.Generate()

	if err != nil {
		log.Fatal(err)
	}

	return document
}
