package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array1 [5]int //você precisa declara o tipo do array, e nenhum outro tipo entra nele, e o TEM QUE TER num,ero de itens, e pode colocar NO MÁXIMO a qtd definida, uma bosta
	fmt.Println(array1)

	// slice
	slice := []int{10,11,12,13,14,15,16}
	fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, 18)
	fmt.Println(slice)



}
