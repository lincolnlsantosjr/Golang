package main

import (
	"fmt"
)

func soma(numeros ...int) (total int) { //função que receberá 1 ou mais inteiros, e retorn um inteiro
	total = 0 //contador
	for _, numero := range numeros { //para indeice (omitido) e numero, percorra os numeros fazendo isso (range numeros)
		total += numero
	}

	return 
}

func main() {
	totalDaSoma := soma(2,2,2,2,2)
	fmt.Println(totalDaSoma)
}
