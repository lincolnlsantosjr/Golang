package main

import "fmt"

func main() {
	// aritmeticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao,divisao, multiplicacao, restoDaDivisao)

	// não pode somer números com tipos diferentes
	var numero1 int8 = 10
	var numero2 int16 = 10
	soma2:= numero1 + int8(numero2)
	fmt.Println(soma2)

	// operadores lógicos

}
