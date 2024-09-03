package auxiliar

import "fmt"

//Escreve uma mensagem na tela falando de qual pacote a mensagem foi escrita
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar!")
	escrever2()
}

 /*Se uma funçao começa com letra maiúscula ela se torna "public" se for com minuscula ela
 só é acessível dentro do próprio pacote ("private") */