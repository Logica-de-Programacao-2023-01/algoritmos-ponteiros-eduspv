package main

//Escreva uma função em Go que receba um ponteiro para um struct Conta
//com campos saldo e titular, e adicione um valor ao saldo da conta.
import "fmt"

type Conta struct {
	Saldo   float64
	Titular string
}

func adicionarValor(ptr *Conta, valor float64) {
	ptr.Saldo += valor
}

func main() {
	conta := Conta{
		Saldo:   100.0,
		Titular: "João",
	}

	ptr := &conta

	fmt.Println("Saldo antes da chamada da função:", conta.Saldo)

	adicionarValor(ptr, 50.0)

	fmt.Println("Saldo após a chamada da função:", conta.Saldo)
}
