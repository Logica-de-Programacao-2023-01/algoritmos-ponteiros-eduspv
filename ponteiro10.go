package main

//Implemente uma função que receba um ponteiro para uma slice e seu tamanho e preencha-o com os n primeiros números primos.
import "fmt"

func isPrimo(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func preencherPrimos(ptr *[]int, tamanho int) {
	slice := *ptr
	i := 2
	for len(slice) < tamanho {
		if isPrimo(i) {
			slice = append(slice, i)
		}
		i++
	}
	*ptr = slice
}

func main() {
	var primos []int
	ptr := &primos
	tamanho := 10

	preencherPrimos(ptr, tamanho)

	fmt.Printf("Os %d primeiros números primos: %v\n", tamanho, primos)
}
