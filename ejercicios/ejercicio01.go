package ejercicios

import (
	"strconv"
)

func ConvertirString(numero string) (int, string) {
	resultado, _ := strconv.Atoi(numero)
	var tamanio string
	if resultado > 100 {
		tamanio = "es mayor que 100"
	} else {
		tamanio = "es menor que 100"
	}

	return resultado, tamanio
}
