package main

import (
	"fmt"
	"math"
)



func main() {

	

	var (
		ladoA      float32
		ladoB      float32 
		hipotenusa float32
	)
	fmt.Println("introduzca dos lados del triangulo")
	fmt.Scan(&ladoA, &ladoB)
	fmt.Println("lado a ", ladoA, "lado b : ", ladoB)

	hipotenusa = float32(math.Pow(float64(ladoA), 2)) + float32(math.Pow(float64(ladoB), 2))
    temp  := (ladoA * ladoB) / 2
	const area float32 = (5 * 6) / 2

	 fmt.Println("Área:", area)
	 fmt.Println("Perímetro:", hipotenusa)
	 fmt.Println("Área:", temp)
}
