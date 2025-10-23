package main

import (
	em "exercise/structs/employees"
	pr "exercise/structs/payroll"
	"fmt"
)

func main() {
	team := []em.Employee{
		{Name: "Анна", Position: "Backend", Skills: []string{"Go", "SQL"}},
		{Name: "Павел", Position: "Data", Skills: []string{"Python", "ETL"}},
		{Name: "Ирина", Position: "QA", Skills: []string{"API testing"}},
	}

	if !team[0].SetBaseSalary(240000) {
		fmt.Println("ошибка: некорректная базовая ставка для", team[0].Name)
	}
	if !team[1].SetBaseSalary(210000) {
		fmt.Println("ошибка: некорректная базовая ставка для", team[1].Name)
	}
	if !team[2].SetBaseSalary(180000) {
		fmt.Println("ошибка: некорректная базовая ставка для", team[2].Name)
	}

	bonusPct := 10.0
	if bonusPct < 0 {
		fmt.Println("предупреждение: отрицательный бонус, устанавливаем 0%")
		bonusPct = 0
	}
	taxPct := 13.0
	if taxPct < 0 || taxPct > 100 {
		fmt.Println("предупреждение: некорректный налог, устанавливаем 13%")
		taxPct = 13
	}

	for _, e := range team {
		gross := pr.CalcGross(e, bonusPct)
		net := pr.CalcNet(gross, taxPct)
		fmt.Printf("%s (%s): gross=%.2f, net=%.2f\n", e.Name, e.Position, gross, net)
		// _ = e.baseSalary // недоступно: приватное поле в пакете employees
	}
}
