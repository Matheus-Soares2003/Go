package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("===========GOROUTINES===========")

	//CONCORRÊNCIA != PARALELISMO

	//PARALELISMO -> DUAS TAREFAS SENDO EXECUTADAS EXATAMENTE AO MESMO TEMPO (CPU COM MAIS DE 1 NUCLEO)
	//CONCORRÊNCIA -> TAREFAS QUE NÃO NECESSARIAMENTE SÃO EXECUTADAS AO MESMO TEMPO | REVEZAMENTO DE TEMPO NAS TAREFAS

	go escrever("Olá Mundo!") // "go" cria uma goroutine
	escrever("Programando em Go!")
	
	//"go" -> execute a função escrever() mas não espera ela terminar de executar para seguir o programa
	//Independente se escrever() terminou, passe para a linha de baixo
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}