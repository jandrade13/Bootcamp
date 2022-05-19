package main

import "fmt"

func getMonth(monthNumber int) string {
	if monthNumber > 12 || monthNumber < 1 {
		return `Valor Invalido!!`
	}
	var months = []string{`Janeiro`, `Fevereiro`, `Março`, `Abril`, `Maio`, `Junho`, `Julho`, `Agosto`, `Setembro`, `Outubro`, `Novembro`, `Dezembro`}
	return months[monthNumber-1]

}

func main() {
	fmt.Println(getMonth(14))
}
