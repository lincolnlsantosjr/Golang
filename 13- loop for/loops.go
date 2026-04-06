package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incremendando: ", i)
		time.Sleep((time.Microsecond))
	}

	// ja declarando variavel junto
	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j: ", j)
		time.Sleep(time.Millisecond)
	}

	// com range, iremos iterar sobre um array -USA BASTANTE

	nomes := []string{"João", "Davi", "Lucas"}

	for indice, nome := range nomes{
		fmt.Println(indice, nome)
	}
}
