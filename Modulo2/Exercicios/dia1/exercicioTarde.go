package main

import "fmt"

func main() {
	palavra := `Palavra`

	for i, word := range palavra {
		fmt.Printf("%d - %s \n", i, string(word))
	}
	fmt.Println("Total de letras: ", len(palavra))

}
