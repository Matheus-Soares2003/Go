package main

import "fmt"

func recuperarExecucao(){
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6") //Mata a execução do programa
	//Quando o panic é executado, antes ele chama todas as funções que tem Defer
}

func main() {
	fmt.Println("==========Cláusula Panic e Recover==========")
	
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós Execução")

	//Não é o melhor jeito de tratar erros em go, mas caso algum pacote tenha uma função que retorne panic
	//você saberá o que fazer
}