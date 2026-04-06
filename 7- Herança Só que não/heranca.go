package main

import "fmt"

type pessoa struct {
	nome   string
	idade  uint
	altura uint
}

type estudante struct {
	pessoa    //aqui estamnos pegando as configurações da pessoa de cima
	curso     string
	faculdade string
}

func main() {

	p1 := pessoa {"joão", 28, 190}

	e1 := estudante{p1, "engenharia", "unisul"}
	fmt.Println(e1.nome)
}
 