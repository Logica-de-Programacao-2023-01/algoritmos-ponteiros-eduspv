package main

import "fmt"

//Escreva uma função que receba um ponteiro para um inteiro e verifique
//se esse inteiro é par ou ímpar. A função deve atualizar o valor do inteiro
//para 0 se ele for par ou para 1 se for ímpar.

func parouimpar(ptr *int) {
	if *ptr%2 == 0 {
		*ptr = 0
	} else {
		*ptr = 1
	}

}
func main() {
	numero := 7
	ptr := &numero

	fmt.Println("Valor antes da chamada da função:", numero)

	parouimpar(ptr)

	fmt.Println("Valor após a chamada da função:", numero)
}
