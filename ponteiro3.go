package main

//Escreva uma função em Go que receba um ponteiro para uma string
//e atualize o valor da string para seu reverso.
import (
	"fmt"
)

func reverterString(ptr *string) {
	runes := []rune(*ptr)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	*ptr = string(runes)
}

func main() {
	texto := "Olá, mundo!"
	ptr := &texto

	fmt.Println("Texto antes da chamada da função:", texto)

	reverterString(ptr)

	fmt.Println("Texto após a chamada da função:", texto)
}
