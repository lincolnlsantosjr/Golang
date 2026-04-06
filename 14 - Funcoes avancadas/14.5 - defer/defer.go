package main

import (
	"fmt"
)

func funcao1() {
	fmt.Println("Executando a função 1")
}
func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool{
	defer fmt.Println("Média calculada. Resultado será retornado!") //tendo defer, executa depois de tudo
	fmt.Println("Entrando na função para verificar se o aluno está aprovado!")

	media := (n1 + n2) / 2

	if media >=6{
		return true
	}
	return false //se não passar no if , retorna false
}

func main(){
	defer funcao1()
	// defer = adiar - ele adia a execução dessa função
	funcao2()
	alunoEstaAprovado(7,8)
}