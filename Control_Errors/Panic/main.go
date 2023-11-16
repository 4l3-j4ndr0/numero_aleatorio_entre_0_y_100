package main

import (
	"fmt"
)

func divide(dividendo,divisor int) {
	validaZero(divisor)
	fmt.Println(divisor/dividendo)
}

// func validaZero(divisor int) {
// 	if(divisor == 0){
// 		// detiene la ejecucion y muestra el error 
// 		panic("No se puede dividir por cero")
// 	}
// }

// si no se establece un mje personalizado igual se activa panic, detiene la 
//ejecucion y muestra un mje predefinido para el tipo de error mostrando 
// lo que paso

// recover se usa para capturar ese error , tratarlo pero no detener la ejecucion 
// del programa siempre y cuando se use tambine defer, ejemplo : 

func validaZero(divisor int) {
	defer func ()  {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if(divisor == 0){
		// detiene la ejecucion y muestra el error 
		panic("No se puede dividir por cero")
	}
}

// OJO : No se recomienda para un uso cotidiano sino solo para situaciones 
// exepcionales o errores graves 

func main() {
	divide(10,20)
	divide(5,2)
	divide(10,0)
	divide(0,10)
	divide(10,2)
	
}