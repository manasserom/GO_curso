package variables

import "fmt"

func MuestroEnteros() {
	var intComun int
	intComun = 0
	intde32 := int32(10)
	intde64 := int64(100)
	fmt.Println("indComun: ", intComun)
	fmt.Println("intde32: ", intde32)
	fmt.Println("intde64: ", intde64)
}
