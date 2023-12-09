package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumeros() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese numero 1:")
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			println(err.Error())
			panic("Dato ingresado es incorrecto")
		}
	}
	fmt.Println("Ingrese numero 2:")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			println(err.Error())
			panic("Dato ingresado es incorrecto")
		}
	}

	fmt.Println("Ingrese leyenda:")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	println(leyenda, numero1*numero2)
}
func ScanNumero() int {
	scanner := bufio.NewScanner(os.Stdin)
	var numero int
	fmt.Println("Ingrese numero:")
	for {
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				println("No ingreso un numero")
				// println(err.Error())
				// panic("Dato ingresado es incorrecto")
			} else {
				break
			}
		}
	}
	return numero
}
