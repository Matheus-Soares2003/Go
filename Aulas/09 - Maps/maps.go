package main

import "fmt"

func main() {

	usuario := map[string]string {
		"nome": "Matheus",
		"sobrenome": "Cabral",
	} //Todas as chaves e valores são apenas do tipo que voce declarou, nesse caso string

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])
	fmt.Println(usuario["sobrenome"])

}