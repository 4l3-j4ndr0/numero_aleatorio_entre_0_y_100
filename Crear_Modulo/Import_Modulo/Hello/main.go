package main

import (
	"fmt"
	"log"

	"github.com/alexroel/greetings"
)

func main() {
	// crea un prefijo que se mostrara en el mje de error 
	log.SetPrefix("greetings : ")
	// con setFlag(0) eliminamos la fecha y la hora que sale por defecto en el mje de error 
	log.SetFlags(0)

	// se llama la funcion Hello del modulo greetings y se guarda su devolucion en la variable message , la funcion hello 
	// del modulo greetings devuelve un mje y un error , por eso es una variable para guardar el mje y otra para el error 
	message, err1 := greetings.Hello("alex")
	ale, err2  := greetings.Hello("Pepe")

	// manejamos el error , en caso de existir , detenemos la app y mostramos el error 
	if err1 != nil || err2 != nil {
		log.Fatal(err1)
		log.Fatal(err2)
	}
	fmt.Println(message)
	fmt.Println(ale)

	fmt.Println("############################################################")

	names := []string {
		"Alex", "Juan", "María", "Lucas","claudia","Alejandra",
	} 

	messages,err := greetings.Hellos(names)
		if err != nil {
			log.Fatal(err)
		}
	fmt.Println(messages)

}