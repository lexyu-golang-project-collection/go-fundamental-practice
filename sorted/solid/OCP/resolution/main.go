package main

import "fmt"

type ReportGenerator interface {
	GenerateReport()
}

type CSVReport struct{}

func (csvRp *CSVReport) GenerateReport() {
	fmt.Println("Generating CSV report")
}

type PDFReport struct{}

func (pdfRp *PDFReport) GenerateReport() {
	fmt.Println("Generating PDF report")
}

type XMLReport struct{}

func (x XMLReport) GenerateReport() {
	fmt.Println("Generating XML report")
}

func generateReports(reports []ReportGenerator) {
	for _, report := range reports {
		report.GenerateReport()
	}
}

func main() {
	reports := []ReportGenerator{
		&CSVReport{},
		&PDFReport{},
		&XMLReport{},
	}

	generateReports(reports)
}
