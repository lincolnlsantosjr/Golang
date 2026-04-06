package main

import "fmt"

func recuperarExecução(){
	if r := recover(); r != nil {
		fmt.Println("execução recuperada com sucesso")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecução()
	media := (n1 + n2) / 2

	if media > 9 {
		return true
	} else if media < 9 {
		return false
	}

	panic("A MEDIA É EXATAMENTE 6")
}

func main() {

	fmt.Println(alunoEstaAprovado(9, 9))
}
