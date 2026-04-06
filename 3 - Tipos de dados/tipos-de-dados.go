package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = -10000000
	fmt.Println(numero)

	var numero2 uint32 = 1000 //uint só aceita numero sem sinal (positivo)
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 123456
	numero0 := 123456
	fmt.Println(numero3)
	fmt.Println(numero0) //funciona, identifica como int, mas ele vai ser baseado na arq do sistema, 32 ou 64

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	// Float, só pode ser 32 ou 64

	var numeroReal1 float32 = 123.5
	var numeroReal2 float64 = 123.50000050

	numeroReal3 := 123.123123 //igual, identifica, mas usa a arq do sistema, 32 ou 64

	fmt.Println(numeroReal1)
	fmt.Println(numeroReal2)
	fmt.Println(numeroReal3)

	// FIM DOS NÚMEROS REAIS

	// STRINGS
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := "A" //teoricamente o go deveria substituir um unico caracter pelo numero dele na tabela ask, mas aqui funcionou como letra normal
	fmt.Println(char)

	// FIM DAS STRINGS
	// Valores zerados

	var texto string //começa zerado (string vazia)
	fmt.Println(texto)
	
	var num1 int //começa zerado (numero vazio)
	fmt.Println(num1)

	var booleano1 bool //comela com false por padrão
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno") //começa com nill por padrão, mas aqui simulei pra ver mesmo
	fmt.Println(erro)
}
