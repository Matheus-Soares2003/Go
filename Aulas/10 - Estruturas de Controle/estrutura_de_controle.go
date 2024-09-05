package main

import "fmt"

func diaDaSemana(num int) string {
	switch num {
	case 1: 
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default: 
		return "Numero Invalido"
	}
}

func diaDaSemana2(num int) string {
	switch {
	case num == 1:
		return "Domingo"
	case num == 2:
		return "Segunda-Feira"
	case num == 3:
		return "Terça-Feira"
	case num == 4:
		return "Quarta-Feira"
	case num == 5:
		return "Quinta-Feira"
	case num == 6:
		return "Sexta-Feira"
	case num == 7:
		return "Sábado"
	default:
		return "Número Inválido"
	}
}

func main() {

	numero := 10

	if numero > 15 {
		fmt.Println("é maior que 15")
	} else {
		fmt.Println("é menor ou igual a 15")
	}

	//if init
	if numero2 := numero; numero2 > 0 { //numero2 recebe numero e logo em seguida é verificado se numero2 é maior que 0
		fmt.Println("Número2 é maior que 0")
	} else if numero2 < 0{
		fmt.Println("Número2 é menor que 0")
	} else {
		fmt.Println("Número2 é igual a 0")
	}

	//fmt.Println(numero2) não vai funcionar pois numero2 é limitada no escopo do if - else

	//SWITCH
	dia := diaDaSemana(6)
	dia2 := diaDaSemana2(2)
	fmt.Println(dia, "e", dia2)
}