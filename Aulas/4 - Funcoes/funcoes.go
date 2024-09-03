package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8){
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {

	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(texto string) {
		fmt.Println(texto)
	}

	f("Texto da função 1")

	resultadosSoma, resultadosSub := calculosMatematicos(30, 20)
	fmt.Println(resultadosSoma, resultadosSub)

	resSoma, _ := calculosMatematicos(15, 5) // _ ignora retorno
	fmt.Println(resSoma)

}