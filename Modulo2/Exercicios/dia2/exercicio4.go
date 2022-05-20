package main

import (
	"errors"
	"fmt"
	"sort"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func minValue(values ...int) int {
	sort.Ints(values)
	return values[0]
}
func maxValue(values ...int) int {
	sort.Ints(values)
	return values[len(values)-1]
}
func averageValue(values ...int) int {
	var sum int
	for _, el := range values {
		sum += el
	}
	return sum / len(values)
}
func operation(typeOperation string) (func(values ...int) int, error) {
	switch typeOperation {
	case minimum:
		return minValue, nil
	case average:
		return averageValue, nil
	case maximum:
		return maxValue, nil
	default:
		return nil, errors.New(`Error`)
	}
}

func main() {

	minhaFunc, err := operation(minimum)
	if err != nil {
		fmt.Println(err)
	}
	averageFunc, err := operation(average)
	if err != nil {
		fmt.Println(err)
	}
	maxFunc, err := operation(maximum)
	if err != nil {
		fmt.Println(err)
	}
	minValue := minhaFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Println(minValue)
	fmt.Println(averageValue)
	fmt.Println(maxValue)
}
