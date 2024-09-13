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

	for opc != 5 {

		fmt.Println("================SISTEMA DE ALUNOS================")
		fmt.Println("\n1 - CADASTRAR ALUNO\n2 - LISTAR ALUNOS\n3 - ATUALIZAR ALUNO\n4 - DELETAR ALUNO\n5 - SAIR")
		fmt.Print("Opção: ")
		fmt.Scan(&opc)
		
		switch opc {
		case 1:
			cadastrarAluno(&alunos)
		case 2:
			listarAlunos(&alunos)
		case 3:
			//atualizarAluno()
		case 4:
			//deletarAluno()
		case 5:
			fmt.Println("Saindo...")
		default:
			fmt.Println("OPÇÃO INVÁLIDA! Tente novamente...")
		}
		fmt.Println("\n====================================================")
	}
}

func cadastrarAluno(listaAlunos *[]Aluno) {

	var aluno Aluno
	
	fmt.Print("RA: ")
	fmt.Scan(&aluno.ra)
	fmt.Print("NOME: ")
	fmt.Scan(&aluno.nome)
	fmt.Print("IDADE: ")
	fmt.Scan(&aluno.idade)
	fmt.Print("CURSO: ")
	fmt.Scan(&aluno.curso)

	*listaAlunos = append(*listaAlunos, aluno)
}

func listarAlunos(listaAlunos *[]Aluno) {
	for _, aluno := range *listaAlunos {
		fmt.Printf("\nRA: %s | Nome: %s | Idade: %d anos | Curso: %s", aluno.ra, aluno.nome, aluno.idade, aluno.curso)
	}
}