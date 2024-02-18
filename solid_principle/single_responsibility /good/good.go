package main

import "fmt"

type EmployeeData struct {
	Name       string
	department string
}

func NewEmployeeData(name, department string) *EmployeeData {
	return &EmployeeData{name, department}
}

type PayCalcrator interface {
	calculatePay(data EmployeeData)
}

type HourReporter interface {
	reportHours(data EmployeeData)
}

type EmployeeRepository interface {
	save()
}
type PayCalcrate struct{}

func NewPayCalcrate() *PayCalcrate {
	return &PayCalcrate{}
}

func (p *PayCalcrate) getRegularHours() {
	fmt.Println("給与計算用の労働計算ロジック")
}

func (p *PayCalcrate) calculatePay(data EmployeeData) {
	p.getRegularHours()
	fmt.Printf("%sの給与を計算しました\n", data.Name)
}

type HourReport struct{}

func NewHourReport() *HourReport {
	return &HourReport{}
}

func (h *HourReport) getRegularHours() {
	// fmt.Println("労働時間レポート専用の労働計算ロジック")
	fmt.Println("労働時間レポート専用の労働計算ロジック 変更したぜ🌟")
}

func (h *HourReport) reportHours(data EmployeeData) {
	h.getRegularHours()
	fmt.Printf("%sの労働時間をレポートしました", data.Name)
}

type EmployeeDB struct{}

func NewEmployeeDB() *EmployeeDB {
	return &EmployeeDB{}
}

func (e *EmployeeDB) save() {

}

func Run() {
	employeeData := NewEmployeeData("takashi", "developer")
	payCalculator := NewPayCalcrate()
	hourReporter := NewHourReport()

	fmt.Println("経理部門")
	payCalculator.calculatePay(*employeeData)
	fmt.Println("")
	fmt.Println("人事部門")
	hourReporter.reportHours(*employeeData)
}
