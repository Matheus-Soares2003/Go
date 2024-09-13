package main

import (
	"fmt"
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
			cadastrarAluno(&alunos)
		case 2:
			//listarAlunos()
		case 3:
			//atualizarAluno()
		case 4:
			//deletarAluno()
		case 5:
			fmt.Println("Saindo...")
		default:
			fmt.Println("OPÇÃO INVÁLIDA! Tente novamente...")
			fmt.Println("====================================================")
		}

		if opc == 5 {
			break
		}
	}
}

func cadastrarAluno(listaAlunos *[]Aluno) {

	var aluno Aluno
	
	fmt.Print("RA: ")
	fmt.Scanf("%d", aluno.ra)
	fmt.Println()
	fmt.Print("NOME: ")
	fmt.Scanf("%s", aluno.nome)
	fmt.Println()
	fmt.Print("IDADE: ")
	fmt.Scanf("%d", aluno.idade)
	fmt.Println()
	fmt.Print("CURSO: ")
	fmt.Scanf("%s", aluno.curso)

	*listaAlunos = append(*listaAlunos, aluno)
}