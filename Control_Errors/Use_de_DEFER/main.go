package main

import (
	"fmt"
)
	// 

func main() {
	// defer --> pasa las sentencias de codigo a una pila , en caso de ser mas de una 
	// la primera en entrar sera la ultima en salir

	// al usar defer lo que se indica es que antes de finalizar la funcion main 
	//se ejecutara la o las lineas de codigo en las que se encuentre defer
	fmt.Println(3)
	fmt.Println(1)
	fmt.Println(2)

	// salida 
	// 3
	// 1
	// 2

	defer fmt.Println(3)
	fmt.Println(1)
	fmt.Println(2)

	// salida
	// 1
	// 2
	// 3
}