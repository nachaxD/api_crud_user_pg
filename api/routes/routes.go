package routes

import (
	"backend/api/handlers"

	"github.com/gorilla/mux"
)

func ConfigureRoutes(r *mux.Router) {
	// allowedOrigins := []string{"http://facturacion.lumonidy.studio", "http://localhost:3000"}

	// c := middleware.CorsMiddleware(allowedOrigins)
	// r.Use(c)

	// r.Handle("/user/get-user", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	handlers.GetUser(w, r)
	// })).Methods("POST")

	r.HandleFunc("/users", handlers.AddUser).Methods("POST")
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.GetUserById).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	// r.Handle("/user/update-user/{UserID}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	handlers.UpdateUser(w, r)
	// })).Methods("PUT")

	// r.Handle("/user/delete-user/{UserID}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	handlers.DeleteUser(w, r)
	// })).Methods("DELETE")
	// Agrega más configuraciones de rutas aquí si es necesario
}
