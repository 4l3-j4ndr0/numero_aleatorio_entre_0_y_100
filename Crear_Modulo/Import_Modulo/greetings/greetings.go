package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {

if name == "" {
	return "", errors.New("nombre vacio")
}

	message := fmt.Sprintf("hola , %v. Bienvenido" , name)
	return message, nil

}
