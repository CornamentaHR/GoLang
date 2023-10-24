package main

import (
	"GoLang/Golang/variables"
	"fmt"
)

func main() {
	estado, texto := variables.ConviertoaTexto(90)
	fmt.Println(estado)
	fmt.Println(texto)
}
