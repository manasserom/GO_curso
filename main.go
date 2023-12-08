package main

import (
	"fmt"

	"github.com/manasserom/GO_curso/ejercicios"
)

func main() {
	numero, mayor := ejercicios.ConvertirString("104")
	fmt.Println(numero)
	fmt.Println(mayor)

	// fmt.Println("Hola Mundo con push automatizado")
	// //variables.MuestroEnteros()
	// //variables.RestoVariables()
	// //estado, text := variables.ConviertoaTexto(123)
	// //os := runtime.GOOS
	// if os := runtime.GOOS; os == "Linux" || os == "OS X." {
	// 	fmt.Println("Esto no es windows, es ", os)
	// } else {
	// 	fmt.Println("Esto es windows: ", os)
	// }

	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Esto es Linux")
	// case "windows":
	// 	fmt.Println("Esto es Windows")
	// case "darwin":
	// 	fmt.Println("Esto es darwin")
	// default:
	// 	fmt.Printf("%s \n", os)
	// }
}
