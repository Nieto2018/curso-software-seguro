package main

import (
	"fmt"
	"io"
	"net/http"
)

// rootHandler es la función que se ejecuta cuando se accede a la ruta /.
// w (ResponseWriter) se usa para escribir la respuesta HTTP.
// r (*Request) contiene la información de la petición HTTP.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// Establecemos el tipo de contenido de la respuesta a texto plano.
	w.Header().Set("Content-Type", "text/plain")

	// Escribimos "Hello, World!" como respuesta.
	fmt.Fprintf(w, "Hello, you are in root!")
}

// helloHandler es la función que se ejecuta cuando se accede a la ruta /hello.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Establecemos el tipo de contenido de la respuesta a texto plano.
	w.Header().Set("Content-Type", "text/plain")

	// Escribimos "Hello, World!" como respuesta.
	fmt.Fprintf(w, "Hello, World!")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received POST request!")

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		fmt.Println("Error reading request body:", err)
		return
	}

	// Imprimimos el cuerpo de la petición en la consola.
	fmt.Print(string(body))

	// Escribimos una respuesta al cliente.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("POST received"))
}

func main() {
	// Definimos un "handler" para la ruta raíz ("/").
	// Cuando alguien acceda a esta ruta, se ejecutará la función rootHandler.
	http.HandleFunc("/", rootHandler)

	// Definimos un "handler" para la ruta raíz ("/hello").
	// Cuando alguien acceda a esta ruta, se ejecutará la función helloHandler.
	http.HandleFunc("/hello", helloHandler)

	http.HandleFunc("/post", postHandler)

	// Iniciamos el servidor en el puerto 8080.
	// La función ListenAndServe bloquea la ejecución y escucha peticiones.
	fmt.Println("Servidor escuchando en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)

	// Si hay un error al iniciar el servidor, lo mostramos.
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}
}
