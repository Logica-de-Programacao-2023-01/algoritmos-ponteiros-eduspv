package main

//Escreva uma função em Go que receba um ponteiro para uma variável float64 e atualize
//o valor da variável para a média aritmética entre o seu valor atual e o valor da constante Pi.
import (
	"fmt"
	"math"
)

func calcularMedia(ptr *float64) {
	pi := math.Pi
	valorAtual := *ptr
	media := (valorAtual + pi) / 2.0
	*ptr = media
}

func main() {
	numero := 3.14
	ptr := &numero

	fmt.Println("Valor antes da chamada da função:", numero)

	calcularMedia(ptr)

	fmt.Println("Valor após a chamada da função:", numero)
}
