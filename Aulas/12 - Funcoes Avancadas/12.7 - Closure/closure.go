package main

import "fmt"

func closure(texto string) func() {
	
	funcao := func() {
		fmt.Println(texto) //Essa variavel texto é a variavel criada dentro da função closure
	}

	return funcao
}

func main() {
	fmt.Println("==========Closure==========")

	//Fuções que referenciam variáveis que estão fora do seu corpo
	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure("Olá Mundo")
	funcaoNova()
}