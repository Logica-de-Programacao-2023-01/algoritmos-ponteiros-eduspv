package main

//Escreva uma função em Go que receba um ponteiro para uma variável
//inteira e atualize o valor da variável para a soma dos valores dos
//seus dois últimos dígitos (por exemplo, se o valor inicial da
//variável for 1234, o novo valor será 3+4=7).
import "fmt"

func somarUltimosDigitos(ptr *int) {
	num := *ptr
	digito1 := num % 10
	num /= 10
	digito2 := num % 10
	soma := digito1 + digito2
	*ptr = soma
}

func main() {
	numero := 1234
	ptr := &numero

	fmt.Println("Valor antes da chamada da função:", numero)

	somarUltimosDigitos(ptr)

	fmt.Println("Valor após a chamada da função:", numero)
}
