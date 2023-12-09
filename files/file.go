package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/manasserom/GO_curso/ejercicios"
)

var filename string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.TablaNumericaGuardada()
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error al crear el archivo: ", err)
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()

}

func AppendTabla() {
	var texto string = ejercicios.TablaNumericaGuardada()
	if !Append(filename, texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(filename string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el Append: ", err)
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteString: ", err)
		return false
	}
	arch.Close()
	return true
}

func LecturaArchivo() {
	// archivo, err := ioutil.ReadFile(filename)
	archivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error al leer archivo: ", err)
		return
	}
	// fmt.Println(string(archivo))
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	archivo.Close()
}
