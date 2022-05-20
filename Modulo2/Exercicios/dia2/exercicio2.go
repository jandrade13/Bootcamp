package main

import (
	"errors"
	"fmt"
)

func average(notas ...float64) (float64, error) {
	var sum float64
	for _, ele := range notas {
		if ele < 0 {
			return 0, errors.New(`Valor Negativo`)
		}
		sum += ele
	}
	return sum / float64(len(notas)), nil
}

func main() {
	result, err := average(2.0, -3.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(`A média é:`, result)
	}
}
