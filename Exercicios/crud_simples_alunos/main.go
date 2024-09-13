package main

import (
	"fmt"

	"golang.org/x/text/cases"
)

type Aluno struct {
	ra string
	nome string
	idade uint
	curso string
	notas []int
}

func main(){

	var (
		alunos []Aluno
		opc int
	)

	fmt.Println("================SISTEMA DE ALUNOS================")
	for {

		fmt.Println("1 - CADASTRAR ALUNO\n2 - LISTAR ALUNOS\n3 - ATUALIZAR ALUNO\n4 - DELETAR ALUNO\n5 - SAIR")
		fmt.Print("Opção: ")
		fmt.Scan(&opc)
		
		switch opc {
		case 1:
			//cadastrarAluno(Aluno)
		case 2:
			//listarAlunos()
		case 3:
			//atualizarAluno()
		case 4:
			//deletarAluno()
		case 5:
			break
		default:
			fmt.Println("OPÇÃO INVÁLIDA! Tente novamente...")
			fmt.Println("====================================================\n\n")
		}
	}
}