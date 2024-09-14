package main

import (
	"fmt"
	"strings"
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
		raPesquisa string
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
			fmt.Print("RA do aluno a ser atualizado: ")
			fmt.Scan(&raPesquisa)
			atualizarRegistroAluno(raPesquisa, &alunos)
		case 4:
			fmt.Print("RA do aluno a ser deletado: ")
			fmt.Scan(&raPesquisa)
			deletarAluno(raPesquisa, &alunos)
		case 5:
			fmt.Println("Saindo...")
		default:
			fmt.Println("OPÇÃO INVÁLIDA! Tente novamente...")
		}
		fmt.Println("\n==================================================")
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

func atualizarRegistroAluno(ra string, listaAlunos *[]Aluno) {

	for idx, aluno := range *listaAlunos {
		if aluno.ra == ra {
			fmt.Printf("\n=====DADOS DO ALUNO=====\nRA: %s | NOME: %s | IDADE: %d anos | CURSO: %s", aluno.ra, aluno.nome, aluno.idade, aluno.curso)
			fmt.Printf("\n\n=====NOVOS DADOS=====\n")

			fmt.Print("NOME: ")
			fmt.Scan(&aluno.nome)
			fmt.Print("IDADE: ")
			fmt.Scan(&aluno.idade)
			fmt.Print("CURSO: ")
			fmt.Scan(&aluno.curso)

			(*listaAlunos)[idx] = aluno
			return
		}
	}

	fmt.Println("ALUNO NÃO ENCONTRADO PARA O RA " + ra)
}

func deletarAluno(ra string, listaAlunos *[]Aluno) {

	for idx, aluno := range *listaAlunos {
		if aluno.ra == ra {
			fmt.Printf("\n=====DADOS DO ALUNO=====\nRA: %s | NOME: %s | IDADE: %d anos | CURSO: %s", aluno.ra, aluno.nome, aluno.idade, aluno.curso)

			var confirmarDeletar string
			fmt.Print("\nDeseja deletar este usuário? [S/N]")
			fmt.Scan(&confirmarDeletar)

			if strings.Compare(confirmarDeletar[:1], "S") == 0 {
				(*listaAlunos)[idx] = Aluno{}
				fmt.Println("Aluno deletado!")
			}
			return
		}
	}

	fmt.Println("ALUNO NÃO ENCONTRADO PARA O RA " + ra)

}