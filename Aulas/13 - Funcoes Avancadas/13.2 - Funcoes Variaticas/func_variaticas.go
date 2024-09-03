package main

import "fmt"


func soma(numeros ...int) (soma int) { //vai receber "n" números inteiros (vira um slice desses "n" números)

	for _, num := range numeros {
		soma += num
	}
	return
}

func escrever(texto string, numeros ...int) { //o param variatico tem que estar em ultimo e só pode ser um por função
	for _, num := range numeros {
		fmt.Println(texto, num)
	}
}

func main() {
	fmt.Println("==========Funções Variáticas==========")
	var soma int = soma()

	fmt.Println(soma)
	escrever("Olá Mundo", 1, 2, 3)
}