package main

import "fmt"

func main(){

	//ARRAYS
	var numeros [5]int

	numeros[0] = 2
	palavras := [2]string{"Matheus", "Giovana"}
	arrDinamico := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //o tamanho fixo do array vai depender de quantos elementos ele possui

	fmt.Println(numeros)
	fmt.Println(palavras)
	fmt.Println(arrDinamico)

	//SLICES
	slice := []int{1, 2, 3, 4, 5, 6} //desse jeito o tamanho do array fica dinamico e ele passa a ser um slice
	fmt.Println(slice)

	slice = append(slice, 7)
	fmt.Println(slice)

	slice2 := arrDinamico[1:5] //Pega uma fatia do array da posição 1 até a posição 4
	fmt.Println(slice2)

	//ARRAYS INTERNOS
	slice3 := make([]int, 5, 10)
	fmt.Println(slice3)

}