package main

import (
	"errors"
	"fmt"
)

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func subtracao(i ...int) int {
	total := i[0]
	for _, v := range i[1:] {
		total -= v
	}
	return total
}

func multiplicacao(i ...int) int {
	total := 1
	for _, v := range i {
		total *= v
	}
	return total
}

func divisao(i ...int) (float64, error) {
	total := i[0]
	for _, v := range i[1:] {
		if v == 0 {
			return float64(0), errors.New("divisão por zero não é permitida")
		}
		total /= v
	}
	return float64(total), nil
}

func main() {
	a := soma(1, 2, 3)
	s := subtracao(30, 20)
	m := multiplicacao(10, 10, 10)
	d, err := divisao(20, 10)
	if err != nil {
		fmt.Println("Erro:", err)
	}

	fmt.Println("Soma:", a)
	fmt.Println("Subtração:", s)
	fmt.Println("Multiplicação:", m)
	fmt.Println("Divisão:", d)

}
