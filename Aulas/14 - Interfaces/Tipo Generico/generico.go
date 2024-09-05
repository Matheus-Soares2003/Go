package main

import "fmt"

//Interfaces vazias funcionam como um parâmetro de tipo genérico (aceitam qualquer tipo)
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	fmt.Println("==========Tipos Genéricos==========")
	generica("Olá Mundo")
	generica(12)
	generica(true)

	mapa := map[string]interface{} {
		"nome": "Matheus Soares",
		"idade": 21,
		"genero": "M",
		"altura": 1.85,
	}

	fmt.Println(mapa)
}