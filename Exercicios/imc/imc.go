package main

import "fmt"

func main() {
	var altura, peso, imc float32

	fmt.Println("==========Descubra seu IMC==========")

	fmt.Print("Digite sua altura: ")
	fmt.Scan(&altura)

	fmt.Print("Digite seu peso: ")
	fmt.Scan(&peso)

	imc = peso / (altura * altura)
	
	fmt.Printf("IMC = %.1f\n", imc)

	switch {
		case imc < 18.5: 	fmt.Println("Abaixo do Peso!")
		case imc < 25: 		fmt.Println("Peso Normal.")
		case imc < 30: 		fmt.Println("Sobrepeso!")
		case imc < 35: 		fmt.Println("Obesidade Grau I!")
		case imc < 40: 		fmt.Println("Obesidade Grau II!")
		case imc >= 40: 	fmt.Println("Obesidade Grau III!")
	}
	
}