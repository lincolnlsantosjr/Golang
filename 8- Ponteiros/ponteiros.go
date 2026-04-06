package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1 //só copiou, quando alterar a variavel1, essa daqui não altera
	fmt.Println(variavel2)
	variavel1++
	fmt.Println(variavel1, variavel2) //add mais 1 na 1, e a 2 não alterou porque era uma cópia

	// ponteiro é uma referencia de memoria
	var variavel3 int //valor default é 0
	var ponteiro *int //valor default é nil

	variavel3 = 100
	// ponteiro = variavel3 - vai dar erro, porque ponteiro não é int, é *int, então só recebe endereços de memória, o certo é assim:
	ponteiro = &variavel3 //feito, aqui recebeu o endereço de memória, e por sua vez, o valor original, se for alterado, altera o ponteiro também
	fmt.Println(variavel3, ponteiro) //veja que aqui o ponteiro deu o endereço de memória do ponteiro, que nesse momento contem a variavel 3
	fmt.Println(variavel3,*ponteiro)//usando o * aqui, nós desrreferenciamos o ponteiro, mostrando o valor que está dentro da memória, nesse caso, da variavell3

	variavel3 = 200
	fmt.Println(variavel3,*ponteiro) //note que não fez cópia, os 2 alteraram, usando ponteiro a gente recebe o valor atualizado (original)
}
