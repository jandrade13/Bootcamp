package main

import "fmt"

func calcTax(salary float64) (result float64) {
	if salary == 50000 {
		result = 50000 * 0.17
	} else if salary == 150000 {
		result = salary * 0.10
	}
	return
}

func main() {
	fmt.Println(calcTax(50000))
	fmt.Println(calcTax(150000))

}
