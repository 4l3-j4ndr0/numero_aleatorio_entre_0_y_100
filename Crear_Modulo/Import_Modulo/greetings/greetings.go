package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {

if name == "" {
	return "", errors.New("nombre vacio")
}
	// message := fmt.Sprintf("hola , %v. Bienvenido" , name)
	message := fmt.Sprintf(RandonFormat(),name)
	return message, nil
}

func RandonFormat() string {
	formats := []string{
		"Hola, %v",
		 "Bienvenido %v, como estas", 
		 "¡Cómo estás %v? Cuanto tiempo sin verte por aca !!"}
	return formats[rand.Intn(len(formats))]
}