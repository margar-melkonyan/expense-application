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

func (s *PDFService) GenDayReport(typeBudget string) core.Document {
	m := getMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err)
	}

	return document
}

func (s *PDFService) GenWeekReport(budgetType string) core.Document {
	m := getMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err)
	}

	return document
}

func (s *PDFService) GenMonthReport(budgetType string) core.Document {
	m := getMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err)
	}

	return document
}

func (s *PDFService) GenYearReport(budgetType string) core.Document {
	m := getMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err)
	}

	return document
}
