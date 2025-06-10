package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Definimos un "handler" para la ruta raíz ("/").
	// Cuando alguien acceda a esta ruta, se ejecutará la función helloHandler.
	http.HandleFunc("/hello", helloHandler)

	// Iniciamos el servidor en el puerto 8080.
	// La función ListenAndServe bloquea la ejecución y escucha peticiones.
	fmt.Println("Servidor escuchando en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)

	// Si hay un error al iniciar el servidor, lo mostramos.
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}
}

// helloHandler es la función que se ejecuta cuando se accede a la ruta raíz.
// w (ResponseWriter) se usa para escribir la respuesta HTTP.
// r (*Request) contiene la información de la petición HTTP.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Establecemos el tipo de contenido de la respuesta a texto plano.
	w.Header().Set("Content-Type", "text/plain")

	// Escribimos "Hello, World!" como respuesta.
	fmt.Fprintf(w, "Hello, World!")
}
