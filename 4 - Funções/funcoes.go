package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// em go, uma função pode ter mais de um retorno
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		return txt
	}

	resultado := f("Texto da função 2")
	fmt.Println(resultado)

	// funcao de calculos acima
	resultadoSoma, resultadoSubtração := calculosMatematicos (10,15)
	fmt.Println(resultadoSoma)
	fmt.Println(resultadoSubtração)
}
