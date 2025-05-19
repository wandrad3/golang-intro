package auxiliar

import "fmt"

func Escrever2() {
	// Esta função não pode ser chamada fora deste pacote
	// porque começa com letra minúscula.
	// Ela é privada ao pacote auxiliar.

	fmt.Println("Escrevendo do pacote auxiliar 2")

}
