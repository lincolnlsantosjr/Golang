package main

import (
	"fmt"
)

func main(){
	fmt.Println("Estruturas de Controle")

	numero := 16

	if numero > 15{
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("menor que 15")
	}

	if outroNumero := numero ; outroNumero > 0 { // if, declara a variável outronumero, atribui o numer 10, e o ; é como se fosse "próxima etapa" e a verificação do if
		fmt.Println("Número é maior que zero")
	} else { //else if também pode por aqui
		fmt.Println("Número é menor que zero")
	}
	// obs: a variavel outronumero criada desse jeito só serve ali dentro
}