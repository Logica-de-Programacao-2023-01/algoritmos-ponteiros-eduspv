package main

//Implemente uma função que receba um ponteiro para uma struct "Livro"
//com campos "Título", "Autor" e "Preço", e que adicione um desconto de 10% no preço do livro.
import "fmt"

type livros struct {
	Titulo string
	Autor  string
	Preco  float64
}

func aplicarDesconto(ptr *livros) {
	desconto := ptr.Preco * 0.1
	ptr.Preco -= desconto
}

func main() {
	livro := livros{
		Titulo: "livros 1",
		Autor:  "Autor 1",
		Preco:  50.0,
	}

	ptr := &livro

	fmt.Println("Preço antes da aplicação do desconto:", livro.Preco)

	aplicarDesconto(ptr)

	fmt.Println("Preço após a aplicação do desconto:", livro.Preco)
}
