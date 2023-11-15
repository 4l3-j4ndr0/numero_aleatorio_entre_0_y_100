package main

import (
	"fmt"
	"math/rand"
)

func main() {
	jugar()
}

func jugar() {
	numAleatorio := rand.Intn(100)
	var numIngresado int
	var intentos int
	const maxIntentos = 10

	for intentos < maxIntentos {
		fmt.Printf("Ingrese un numero (Intentos restantes : %d) \n", maxIntentos-intentos)
		fmt.Scanln(&numIngresado)
		if numIngresado == numAleatorio {
			fmt.Printf("Felicidades , acertaste el numero aleatorio (%d) \n", numIngresado)
			jugarNuevamente()
			return
		} else if numIngresado < numAleatorio {
			fmt.Printf("El numero a adivinar es mayor que el ingresado (%d) \n", numIngresado)
		} else if numIngresado > numAleatorio {
			fmt.Printf("El numero a adivinar es menor que el ingresado (%d) \n", numIngresado)
		}
		fmt.Println(numAleatorio)
		intentos++
	}
	fmt.Printf("Lo siento , has perdido. El numero era %d", numAleatorio)
	jugarNuevamente()
}

func jugarNuevamente() {
	var eleccion string
	fmt.Println("Quieres jugar nuevamente ? (s/n) :")
	fmt.Scanln(&eleccion)

	switch eleccion {
	case "s":
		jugar()
	case "n":
		fmt.Println("Gracias por jugar")
	default:
		fmt.Println("Eleccion invalida, favor vuelvela a ingresar")
		jugarNuevamente()
	}

	
}
