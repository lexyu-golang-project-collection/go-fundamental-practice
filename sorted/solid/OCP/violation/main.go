package main

import "fmt"

type Report struct{}

func (r Report) GenerateReport(reportType string) {
	if reportType == "CSV" {
		fmt.Println("Generating CSV report")
	} else if reportType == "PDF" {
		fmt.Println("Generating PDF report")
	} // 實作新功能，需要修改既有代碼，違反OCP
}

func main() {
	report := &Report{}
	report.GenerateReport("CSV")
	report.GenerateReport("PDF")
}
