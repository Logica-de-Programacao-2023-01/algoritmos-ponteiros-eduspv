package main

import "fmt"

// Escreva uma função que receba um ponteiro para um inteiro e um valor inteiro n.
// A função deve atualizar o valor do inteiro para a soma dos n primeiros números naturais.
func numerosNaturais(p1 *int, n1 int) {
	soma := (n1 * (n1 + 1)) / 2
	*p1 = soma
}
func main() {
	var p1 int
	n1 := 5
	numerosNaturais(&p1, n1)
	fmt.Println(p1)

}
