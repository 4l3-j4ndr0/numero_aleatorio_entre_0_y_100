package main

import (
	"errors"
	"fmt"
	"strconv"
)
	// con el pakete error se puede manejar el error , el cual recibe un string que seria el 
	// mensaje personalizado del error 
func divide (dividendo, divisor int)(int, error){
	if(dividendo == 0){
		return 0,errors.New("No se puede dividir por cero")
	}
	return dividendo / divisor, nil
}

func main() {
	str := "123"
	num,err := strconv.Atoi(str)
	if err != nil{
		fmt.Println("Error : ",err)
	}else{
		println(num)
	}

	fmt.Println(divide(0,1))
}