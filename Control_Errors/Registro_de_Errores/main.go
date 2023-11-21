package main

import (
	"log"
	"os"
)

func main() {
	// // se usa para agregar un prefijo a los mjes de registros del programa 
	// log.SetPrefix("main( PREFIJO )")

	// // emite un log ( fecha, hora y mje colocado ) y detiene el programa 
	// log.Fatal("meje de registro FATAL")
	// // esta linea ya no se ejecuta
	// log.Println("meje de registro 2")

	// Abrimos o creamos un archivo, el 0644 es que el archivo tendra permiso de escritura y lectura para el propietario del archivo
	// os.O_CREATE --> crea el archivo
	// os.O_APPEND --> abre el archivo
	// os.O_WRONLY --> establece el permiso de solo escritura
	// 0644 --> valor octal para los permisos del archivo, en este caso se estableceran los permiso de lectura y escritura del archivo solo para 
	// el propietario  del archivo 
	file,error := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
	if error != nil{
		log.Fatalln("Error al abrir archivo: ",error)
	}
	// crea un flujo para guardar la informacion del log en en el archivoantes creado (info.log)
	log.SetOutput(file)
	// el mje que se guardara en el archivo antes creado 
	log.Print("oye soy un log")


	// se debe cerrar el flujo abierto , por lo que usamos el defer para q se ecute al final 
	defer file.Close()
}