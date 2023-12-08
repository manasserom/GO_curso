package main

import (
	"fmt"

	"github.com/manasserom/GO_curso/variables"
)

func main() {
	fmt.Println("Hola Mundo con push automatizado")
	//variables.MuestroEnteros()
	//variables.RestoVariables()
	estado, text := variables.ConviertoaTexto(123)
	fmt.Println(estado)
	fmt.Println(text)
}
