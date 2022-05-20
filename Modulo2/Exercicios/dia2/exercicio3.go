package main

import (
	"errors"
	"fmt"
)

func categoryA(hours int) int {
	if hours > 160 {
		result := 4500.00 * hours
		return result
	} else {
		return 3000 * hours
	}
}
func categoryB(hours int) int {
	if hours > 160 {
		result := 1500 * 1.2 * hours
		return result
	} else {
		return 1500 * hours
	}
}

func categoryC(hours int) int {
	return 1000 * hours

}

func calcSalary(category string, totalHours int) (int, error) {
	switch category {
	case `A`:
		return categoryA(totalHours), nil
	case `B`:
		return categoryB(totalHours), nil
	case `C`:
		return categoryC(totalHours), nil
	default:
		return 0, errors.New(`Categoria Invalida`)
	}
}
func main() {
	result, err := calcSalary(`C`, 162)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	result, err = calcSalary(`B`, 176)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	result, err = calcSalary(`A`, 172)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
