package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	employeesMore21 := 0
	for _, idade := range employees {
		if idade > 21 {
			employeesMore21++
		}

	}
	fmt.Println(`Quantidade maior que 21 anos:`, employeesMore21)
	employees[`Carlos`] = 25
	fmt.Println(employees)
	delete(employees, `Pedro`)
	fmt.Println(employees)
}
