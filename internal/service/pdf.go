package service

import (
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"log"
)

type PDFService struct {
}

func NewPdfService() *PDFService {
	return &PDFService{}
}

func getMaroto() core.Maroto {
	m := maroto.New()

	return m
}

func (s *PDFService) GenDayReport(budgetType string) {
	m := getMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err)
	}

	err = document.Save("docs/assets/pdf/simplestv2.pdf")
	if err != nil {
		log.Fatal(err)
	}
}

func (s *PDFService) GenWeekReport(budgetType string) {

}

func (s *PDFService) GenMonthReport(budgetType string) {

}

func (s *PDFService) GenYearReport(budgetType string) {

}
