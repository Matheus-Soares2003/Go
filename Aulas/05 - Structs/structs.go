package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero int8
}

func main() {
	fmt.Println("Arquivos structs")

	var u usuario
	u.nome = "Matheus"
	u.idade = 21
	fmt.Println(u)

	endereco := endereco{"P. Sherman - 42, Wallaby way, Sidney", 42}

	usuario2 := usuario{"Giovana", 22, endereco}	
	fmt.Println(usuario2)
}