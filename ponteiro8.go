package main

//Crie uma função que receba um ponteiro para uma variável como argumento e modifique o valor da variável dentro da função.
//Certifique-se de que o ponteiro aponte para uma área de memória válida e que a memória seja liberada após o uso.
import (
	"fmt"
	"unsafe"
)

func modificarValor(ptr *int) {
	*ptr = 42 // Modifica o valor da variável apontada pelo ponteiro
}

func main() {
	numero := 10
	ptr := &numero

	fmt.Println("Valor antes da chamada da função:", numero)

	modificarValor(ptr)

	fmt.Println("Valor após a chamada da função:", numero)

	// Libera a memória após o uso
	ptr = nil
	size := unsafe.Sizeof(*ptr)
	ptr = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + size))
}
