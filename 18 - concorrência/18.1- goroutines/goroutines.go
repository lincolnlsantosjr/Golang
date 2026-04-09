package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRENCIA != PARALELISMO

	go escrever ("Olá, Mundo!") //goroutine - ela faz meio que "começa a executar esse código, mas não espera ele terminar pra continuar o código"
	escrever("Programando em Go!")
}

func escrever(texto string){
	for{
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
