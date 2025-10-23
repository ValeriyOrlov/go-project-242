package payroll

import (
	employees "exercise/structs/employees"
)

func CalcGross(e employees.Employee, bonusPct float64) float64 {
	base := e.BaseSalary()
	return base + base*bonusPct/100
}

func CalcNet(gross float64, taxPct float64) float64 {
	return gross * (1 - taxPct/100)
}
