package main

import "fmt"

func main() {

	var (
		temp    float64 = 35.2
		umidade float64 = 0.20
		press   float64 = 1.00
	)
	fmt.Printf(`Tem: %.2f C, umidade:%.2f%%, Pressão: %.2f Atm`, temp, umidade*100, press)
}
