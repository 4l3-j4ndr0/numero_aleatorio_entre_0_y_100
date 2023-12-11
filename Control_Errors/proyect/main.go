package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// estructura de contactos
type Contact struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// guarda contactos en un json
func saveContactToFile(contact []Contact) error {
	// validamos q no exista un error 
	file,err := os.Create("contact.json")
	if err!= nil {
		return err
	}

	// cerramos el flujo usando defer para que se ejecute al final de todo 
	defer file.Close()

	// encoder se usa para convertir estructuras de datos en archivos json
	encoder := json.NewEncoder(file)

	// se usa para codificar el slice en un formato json y escribirlo en un archivo , se conoce como codificacion o serializacion de datos
	err = encoder.Encode(contact)

	if err != nil {
		return err
	}

	return nil
}


// cargar contactos del archivo json
func loadContactFromFile(contacts *[]Contact )error{
	file, err := os.Open("contact.json")
	if err != nil {
		return  err
	}	
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
 return nil
}

func main() {
	// slice de contactos
	var contacts []Contact

	// cargar contactos existentes 
	err := loadContactFromFile(&contacts)
	if err != nil {
		fmt.Println("error al cargar los contactos", err)
	}

	// mostrar opciones al usuario
	fmt.Println("***** Toma una opcion ***** \n",
	"1. Agregar contacto \n",
	"2. Mostrar todos los contactos \n",
	"3. Salir" )
	
	// leer opcion del usuario 
	var opcion int
	fmt.Scanln(&opcion)
	switch opcion {
		case 1: saveContactToFile(contacts)
		case 2: loadContactFromFile(&contacts)
		case 3: os.Exit(0)
		default: fmt.Println("opcion no valida")
		
		
	}
}