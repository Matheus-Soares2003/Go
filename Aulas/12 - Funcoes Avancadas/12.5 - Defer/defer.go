package main

import "fmt"

func funcao1(){
	fmt.Println("Executando a função 1")
}

func funcao2(){
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {

	defer fmt.Println("Média calculada! Resultado será retornado") //vai ser executado imediatamente antes do return
	fmt.Println("Verificando se aluno está aprovado")
	media := (n1 + n2) / 2

	return media >= 6
}

func main() {
	fmt.Println("==========Cláusula Defer==========")

	/*DEFER = ADIAR
	defer funcao1() //Vai adiar a função até o ultimo momento possivel
	funcao2()
	fmt.Println("Olá Mundo!")*/


	fmt.Println(alunoEstaAprovado(2, 8))

	//Muito utilizado para fechar conexao com banco de dados já que não existe try catch em golang
	
}