package main

import (
	"fmt"
)

func main(){
	usuario := map[string]string{ //dentro do colchetes é o tipo das chaves, fora dos colchetes, os valores
		"nome": "pedro",
		"sobrenome": "silva",
	}
// muito ruim de usar
	fmt.Println((usuario["nome"]))
}