package main

import "fmt"

type pessoa struct {
	nome  string
	idade uint8
}

type estudante struct {
	pessoa    //"Herança" -> pega todos os campos de pessoa e copia para estudante
	curso     string
	faculdade string
}

func main() {

	pessoa := pessoa{"Matheus Cabral", 21}
	estudante := estudante{pessoa, "Ciência da Computação", "UNINOVE"}

	fmt.Println(estudante.nome)
	fmt.Println(estudante.curso)
	fmt.Println(estudante.faculdade)

}