package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("===========WAITGROUP===========")

	var waitGroup sync.WaitGroup
	waitGroup.Add(2) //Grupo de espera com 2 goroutines

	go func() {
		escrever("Olá Mundo")
		waitGroup.Done() //sinaliza que a função escrever terminou sua execução
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done()
	}()

	waitGroup.Wait() //Espera todas as goroutines acabarem
}

func escrever(texto string) {
	for i := 0; i < 5; i++{
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}