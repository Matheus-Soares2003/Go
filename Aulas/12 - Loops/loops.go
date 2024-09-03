package main

import (
	"fmt"
	//"time"
)

func main() {
	fmt.Println("==========LOOPS==========")

	//Em go só temos o loop for que pode ser usado de diferentes maneiras

	//Parecido com WHILE
	cont := 0
	for cont <= 5 {
		//time.Sleep(time.Second) pausa por 1 segundo
		fmt.Printf("%d ", cont)
		cont++
	}

	fmt.Println()

	//For loop clássico
	for i := 0; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println()

	//For com a cláusa range (usado para iterar sobre um array ou slice ou string)
	nomes := [3]string{"Matheus", "Giovana", "Fernando"}

	for indice, nome := range nomes { //o primiero parametro é sempre o indice e o segundo é o item do array
		fmt.Printf("%d - %s\n", indice, nome)
	}

	//Caso não queira o indice você pode ignora-lo usando o "_"
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	fmt.Println()

	//Com string se voce usar o range ele vai te retornar o numero de ASCII daquela letra
	for idx, letra := range "PALVARA" {
		fmt.Println(idx, "-", letra)
	}

	//Caso queira a letra em si é so fazer o cast para string
	for idx, letra := range "PALVARA" {
		fmt.Println(idx, "-", string(letra))
	}

	fmt.Println()

	//Também é possível iterar por um map
	usuario := map[string]string {
		"nome": "Yuri",
		"sobrenome": "Alberto",
	}

	for chave, valor := range usuario {
		fmt.Printf("%s => %s\n", chave, valor)
	}
	//Não é possível usar range em Structs


	//Loop infinito
	numero := 0
	for {
		fmt.Println(numero)
		numero++
		if numero == 15 {
			break //Só é possivel parar o loop com o break ou lançando uma exceção (não recomendavel)
		}
	}
}