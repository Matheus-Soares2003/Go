package main

import "fmt"

func fibonnaci(posicao uint) uint {

	if posicao <= 1 {
		return posicao
	}

	return fibonnaci(posicao - 2) + fibonnaci(posicao - 1)
}

func main() {
	fmt.Println("==========Funções Recursivas==========")

	posicao := uint(3)
	fmt.Println(fibonnaci(posicao))
}