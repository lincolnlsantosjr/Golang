package main

import (
	"fmt"
)

func main(){
	retorno := func(texto string)string{ //variavel retorno recebe uma função anonima
		return fmt.Sprintf("recebido: %s", texto)
	}("parametro passado")//aqui a função ja é executada logo após encerramwento (default de função anonima)

	fmt.Println(retorno)
}
