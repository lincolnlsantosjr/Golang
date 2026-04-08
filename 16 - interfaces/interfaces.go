package main

import "fmt"

type forma interface {
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 { //método implementando pra cada retangulo a área
	return r.altura * r.largura
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f", f.area()) //função da interface
}

func main() {
	r := retangulo{10, 14}
	escreverArea(r)
}
