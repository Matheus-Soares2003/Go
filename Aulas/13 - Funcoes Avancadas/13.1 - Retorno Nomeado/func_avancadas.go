package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) /*retorno nomeado*/ {
	soma = n1 + n2
	subtracao = n1 - n2
	return //vai retornar as variaveis declaradas nos parenteses acima (soma e subtração)
}

func main() {
	fmt.Println("==========Retorno Nomeado==========")
	som, sub := calculosMatematicos(5, 3)

	fmt.Printf("Soma: %d\nSubtração: %d", som, sub)
}