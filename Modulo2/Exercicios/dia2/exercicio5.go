package main

import (
	"errors"
	"fmt"
)

const (
	dog   = "dog"
	cat   = "cat"
	ham   = "ham"
	taran = "taran"
)

func animaldog(qtdAnimals int) float64 {
	return float64(qtdAnimals) * 10.00
}

func animalcat(qtdAnimals int) float64 {
	return float64(qtdAnimals) * 5.00
}

func animalham(qtdAnimals int) float64 {
	return float64(qtdAnimals) * 0.250
}

func animaltaran(qtdAnimals int) float64 {
	return float64(qtdAnimals) * 0.150
}

func Animal(nameAnimal string) (func(qtdAnimals int) float64, error) {
	switch nameAnimal {
	case dog:
		return animaldog, nil
	case cat:
		return animalcat, nil
	case ham:
		return animalham, nil
	case taran:
		return animaltaran, nil
	default:
		return nil, errors.New(`Error`)

	}
}
func main() {
	var amount float64
	animalDog, msg := Animal(dog)

	if msg != nil {
		fmt.Println(msg)
	} else {
		fmt.Println(animalDog(5))
	}

	animalCat, msg := Animal(cat)
	if msg != nil {
		fmt.Println(msg)
	} else {
		fmt.Println(animalCat(8))
	}

	amount += animaldog(5)
	amount += animalCat(8)

	fmt.Println(amount)
}
