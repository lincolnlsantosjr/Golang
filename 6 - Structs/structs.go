package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	rua string
	num uint
}

func main() {
	fmt.Println("Arquivo structs")
	vendereco1 := endereco{"rua irma bonavita", 1505}
	var u usuario
	u.nome = "Davi"
	u.idade = 21

	usuario2 := usuario{"davi", 21,vendereco1 }
	fmt.Println(u)
	fmt.Println(usuario2)

	fmt.Println(usuario2.endereco)
}
