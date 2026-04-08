package main

import (
	"fmt"
)

// função executada antes da função main, SÓ PODE TER UMA POR ARQUIVO
func init(){
	fmt.Println("executando a função init")
}

func main() {
	fmt.Println("Init")

}
