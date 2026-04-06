package main

import (
	"fmt"
)

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	// não usa := pois ja foram decalaradas no titulo da função
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	subtracao, soma := calculosMatematicos(10, 20)
	fmt.Println(subtracao,soma)
}
