package main

import (
	"fmt"
)

func inverterSinaisComPonteiros (numero *int){
	*numero = *numero * -1
}

func main() {
	fmt.Println("oi")

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinaisComPonteiros(&novoNumero) //aqui passamos o endereço de memoria do novoNumero, alterando assim o valor originl la de dentro, não cópia
	fmt.Println(novoNumero)
	
}
