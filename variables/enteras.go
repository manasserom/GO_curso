package variables

import (
	"fmt"
	"strconv"
)

func MuestroEnteros() {
	var intComun int
	intComun = 0
	intde32 := int32(10)
	intde64 := int64(100)
	fmt.Println("indComun: ", intComun)
	fmt.Println("intde32: ", intde32)
	fmt.Println("intde64: ", intde64)
}
func ConviertoaTexto(numero int) (bool, string) {
	var texto string
	texto = strconv.Itoa(numero)
	return true, "t" + texto + "t"
}
