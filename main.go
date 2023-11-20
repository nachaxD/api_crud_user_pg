package main

import (
	"fmt"
	"log"
	"net/http"

	"backend/api/routes"

	"github.com/gorilla/mux"
)

func main() {
	// Crear un enrutador utilizando gorilla/mux
	r := mux.NewRouter()

	// Configurar rutas desde routes/routes.go
	routes.ConfigureRoutes(r)

	// Puerto en el que escuchará el servidor
	port := ":8080"

	// Iniciar el servidor web y configurar para escuchar en el puerto
	fmt.Printf("Servidor web iniciado y escuchando en http://localhost%s...\n", port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v\n", err)
	}
}
