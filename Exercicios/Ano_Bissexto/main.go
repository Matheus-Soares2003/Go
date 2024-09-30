package main

import "fmt"

func ano_bissexto(ano int) bool {

	//Deixei as condições em variaveis para facilitar a leitura do código
	divisivelPor4 := (ano%4 == 0)
	divisivelPor100 := (ano%100 == 0)
	divisivelPor400 := (ano%400 == 0)

	return (divisivelPor4 && !divisivelPor100) || (divisivelPor4 && divisivelPor100 && divisivelPor400)
}

func main() {

	anos_teste := []int{2100, 2024, 1800, 2016, 2000, 1600}

	for _, ano := range anos_teste {
		fmt.Println(ano, " é bissexto?\t", ano_bissexto(ano))
	}

}