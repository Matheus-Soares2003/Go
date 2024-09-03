package main

import (
	"errors"
	"fmt"
)

func main() {
	var (
		numeroInteiro int = 100
		numeroReal float64 = 9.5
		palavra string = "Matheus Soares"
	)
	
	fmt.Println(numeroInteiro, numeroReal)
	fmt.Println(palavra)

	//alias | apelido
	//rune = int32, muito utilizado quando trabalhamos com números que representam caracteres (ASCII)
	var num rune = 123456 
	fmt.Println(num)

	//GoLang não possui o tipo "char", ele vira um número que representa o caracter na tabela ASCII
	caracter := 'A'
	fmt.Println(caracter)
	
	//Valor quando não se inicializa uma variável
	//Uma string se inicializa vazia ("") um inteiro se inicializa com 0, boolean com false, etc..
	var texto string
	var numero int
	fmt.Println(texto, numero)

	//Tipo de Dado de erro (se inicializa com nil)
	var erro error
	fmt.Println(erro)

	//Para se inicializar uma variável desse tipo é necessário usar um pacote interno chamado "errors"
	var erro2 error = errors.New("Erro interno!")
	fmt.Println(erro2)
}