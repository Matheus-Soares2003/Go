package main

import "fmt"


func main() {
	fmt.Println("==========Funções Anonimas==========")

	func() {
		fmt.Println("Olá Mundo!")
	}() //Por ter esse "()" no final, essa função será executada automaticamente

	func(texto string) {
		fmt.Println(texto, "Recebido")
	}("Escrevendo texto")

	//É possível armazenar funções dentro de variaveis e depois invoca-las como uma função
	escreveOlaMundo := func(texto string) {
		fmt.Println(texto)
	}

	escreveOlaMundo("Escrevendo um novo texto!")

	soma := func(n1, n2 int) int {
		return n1 + n2
	}(5, 4)

	soma2 := func(n1, n2 int) int {
		return n1 + n2
	}

	fmt.Println(soma)
	fmt.Println(soma2(3, 3))

	
}