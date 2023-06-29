package main

//Escreva uma função em Go que receba um ponteiro para um struct livros com campos título e autor,
//e altere o título do livro para "Desconhecido" se o autor for "Anônimo".
import "fmt"

type Livro struct {
	Titulo string
	Autor  string
}

func alterarTitulo(ptr *Livro) {
	if ptr.Autor == "Anônimo" {
		ptr.Titulo = "Desconhecido"
	}
}

func main() {
	livro := Livro{
		Titulo: "livros 1",
		Autor:  "Anônimo",
	}

	ptr := &livro

	fmt.Println("Título antes da chamada da função:", livro.Titulo)

	alterarTitulo(ptr)

	fmt.Println("Título após a chamada da função:", livro.Titulo)
}
