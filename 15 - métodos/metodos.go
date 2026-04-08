package main

import (
	"fmt"
)

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() { //é um método, diz que cada usuário (da struct) receberá o método salvar
	fmt.Printf("salvando os dados do Usuário %s no banco de dados\n", u.nome)
}

// com ponteiro
func (u *usuario) fazerAniversario(){
	u.idade++
}

func main() {
	usuario1 := usuario{"Lincoln", 30}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Daniel", 30}
	fmt.Println(usuario2.nome)
	usuario2.salvar()

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)

}
