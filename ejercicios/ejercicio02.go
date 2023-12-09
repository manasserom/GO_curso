package ejercicios

import (
	"fmt"

	"github.com/manasserom/GO_curso/teclado"
)

func TablaNumerica() {
	var numero int = teclado.ScanNumero()
	fmt.Println("TABLA DEL ", numero)
	for i := 1; i < 11; i++ {
		println(i * numero)
	}
}
