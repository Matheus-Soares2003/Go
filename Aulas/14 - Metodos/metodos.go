package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
}

//Métodos da struct usuario
func (u usuario) apresentar() {
	fmt.Printf("Olá! Me chamo %s e tenho %d anos\n", u.nome, u.idade)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

//Nesse caso fazendo uma referencia a struct com ponteiros para alterar valores
func (u *usuario) fazerAniversario() {
	u.idade++
}


func main() {
	fmt.Println("==========Métodos==========")
	usuario1 := usuario{"Matheus", 21}

	usuario2 := usuario{"Miguel", 18}
	usuario2.fazerAniversario()

	usuario1.apresentar()
	usuario2.apresentar()
}