package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler) // Define el handler para la ruta principal

	port := ":8080" // Define el puerto del servidor
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)

	// Inicia el servidor
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Error iniciando el servidor: %s\n", err)
	}
}
