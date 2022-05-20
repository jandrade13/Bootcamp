package main

import "fmt"

type info struct {
	salario       float64
	idade         int
	anosAtividade int
	trabalhando   bool
}

func checkLoan(user info) string {
	if user.salario < 100000 {
		return (`Salario baixo`)
	} else if user.idade < 22 {
		return (`Idade Baixa`)
	} else if user.anosAtividade < 1 {
		return (`Anos de Atividade Menor que 1`)
	} else if !user.trabalhando {
		return (`Desempregado`)
	} else {
		return (`Parabens, tu deve o banco!!!`)
	}
}

func main() {
	pessoa1 := info{
		salario:       100000.00,
		idade:         25,
		anosAtividade: 2,
		trabalhando:   true,
	}
	result := checkLoan((pessoa1))
	fmt.Println(result)
}
