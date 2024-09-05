package main

import "fmt"

func main() {

	var nome1 string = "Matheus Soares"
	nome2 := "Galvão Bueno"

	var (
		nome3 string = "Maria"
		nome4 string = "Giovana"
		num1 int = 2
		num2 int = 5
	)

	numero1, numero2 := 5, 10

	const pi float32 = 3.1415

	const (
		gravidade float32 = 10
		euler float32 = 2.71828
	)

	fmt.Println(nome1)
	fmt.Println(nome2)

	fmt.Println(num1)
	fmt.Println(num2)

	fmt.Println(nome3)
	fmt.Println(nome4)

	fmt.Println(numero1, numero2)

	numero1, numero2 = numero2, numero1

	fmt.Println(numero1, numero2)

}