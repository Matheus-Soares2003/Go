package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1 //variacvel2 é uma copia de variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++ //se alterarmos a variavel1 ela não altera a variável2
	fmt.Println(variavel1, variavel2)

	//PONTEIRO É UMA REFERENCIA DE MEMORIA
	var variavel3 int
	var ponteiro *int //Guarda um endereço de memória de um valor int

	variavel3 = 100
	ponteiro = &variavel3 //"&" devolve o endereco da variavel

	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro)

	variavel3 += 100
	fmt.Println(variavel3, *ponteiro)
}