package main

import "fmt"

//Esse parâmetro recebe uma cópia do valor informado
func inverterSinal(num int) int {
	return num * -1
}

//Nesse caso não precisamos de retorno pois iremos alterar diretamente o valor da variável
func inverterSinalPonteiro(num *int) {
	*num = *num * -1
}

func mudarTexto(txt *string){
	*txt = "Texto Alterado"
}

func mudarArr(lista *[]int) {
	*lista = append(*lista, 10)
}

func main() {
	fmt.Println("==========Funções com ponteiros==========")

	numero := 20
	numeroSinalInvertido := inverterSinal(numero) //Cópia do valor 20
	fmt.Println(numeroSinalInvertido)

	numero2 := 10
	inverterSinalPonteiro(&numero2) //é passado o endereço de memória da variavel
	fmt.Println(numero2)

	texto := "Olá Mundo!"
	mudarTexto(&texto)
	fmt.Println(texto)

	numeros := []int{1, 2, 3}
	mudarArr(&numeros)
	fmt.Println(numeros)
}