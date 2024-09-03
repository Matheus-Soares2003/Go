package main

import (
	"modulo/auxiliar"
	"fmt"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo no arquivo main!")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("matheuscabral@gmail.com")

	fmt.Println(erro)
}