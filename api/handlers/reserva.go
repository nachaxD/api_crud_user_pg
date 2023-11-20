package handlers

import (
	"backend/api/models" // Asegúrate de tener una estructura Reserva definida en tu paquete models
	"backend/api/utils"  // Asegúrate de importar tu paquete utils con la función OpenDB
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func GetReservasPorUsuario(w http.ResponseWriter, r *http.Request) {
	var reservas []models.Reserva // Asegúrate de tener una estructura Reserva definida en tu paquete models

	// Lee el cuerpo de la solicitud
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	// Convierte el cuerpo JSON a una estructura que contenga el campo "email"
	var requestBody map[string]string
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		http.Error(w, "Error al analizar el cuerpo JSON", http.StatusBadRequest)
		return
	}

	email, ok := requestBody["email"]
	if !ok {
		http.Error(w, "El campo 'email' es requerido en el cuerpo JSON", http.StatusBadRequest)
		return
	}

	// Abre la conexión a la base de datos
	db, err := utils.OpenDB()
	if err != nil {
		http.Error(w, "Error al conectar a la base de datos", http.StatusInternalServerError)
		return
	}

	// Utiliza el método Where para filtrar por id_usuario
	db.Where("id_usuario = ?", email).Find(&reservas)

	if len(reservas) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Configura la respuesta JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&reservas)
}

func GetReservasPorEstado(w http.ResponseWriter, r *http.Request) {
	var reservas []models.Reserva

	// Lee el cuerpo de la solicitud
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	// Convierte el cuerpo JSON a una estructura que contenga los campos "email" y "estado"
	var requestBody map[string]string
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		http.Error(w, "Error al analizar el cuerpo JSON", http.StatusBadRequest)
		return
	}

	email, ok := requestBody["email"]
	if !ok {
		http.Error(w, "El campo 'email' es requerido en el cuerpo JSON", http.StatusBadRequest)
		return
	}

	estado, ok := requestBody["estado"]
	if !ok {
		http.Error(w, "El campo 'estado' es requerido en el cuerpo JSON", http.StatusBadRequest)
		return
	}

	// Abre la conexión a la base de datos
	db, err := utils.OpenDB()
	if err != nil {
		http.Error(w, "Error al conectar a la base de datos", http.StatusInternalServerError)
		return
	}

	// Utiliza el método Where para agregar condiciones a la consulta
	db.Where("id_usuario = ? AND estado = ?", email, estado).Find(&reservas)

	if len(reservas) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Configura la respuesta JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&reservas)
}

func BorrarReserva(w http.ResponseWriter, r *http.Request) {
	// Lee el cuerpo de la solicitud
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	// Convierte el cuerpo JSON a una estructura que contenga el campo "id"
	var requestBody map[string]interface{}
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		http.Error(w, "Error al analizar el cuerpo JSON", http.StatusBadRequest)
		return
	}

	// Convierte el id a un valor numérico
	id, ok := requestBody["id"].(float64)
	if !ok {
		http.Error(w, "El campo 'id' es requerido en el cuerpo JSON y debe ser un número", http.StatusBadRequest)
		return
	}

	// Convierte el id a un entero sin decimales
	idInt := int(id)

	// Abre la conexión a la base de datos
	db, err := utils.OpenDB()
	if err != nil {
		http.Error(w, "Error al conectar a la base de datos", http.StatusInternalServerError)
		return
	}

	// Busca la reserva por ID
	var reserva models.Reserva
	db.First(&reserva, idInt)

	if reserva.ID == 0 {
		http.Error(w, "Reserva no encontrada", http.StatusNotFound)
		return
	}

	// Verifica que la reserva esté en estado 'Aprobado'
	if reserva.Estado != "Aprobado" {
		http.Error(w, "No se puede borrar la reserva, no está en estado 'Aprobado'", http.StatusForbidden)
		return
	}

	// Calcula la diferencia de días entre la fecha actual y la fecha_reserva
	diferenciaDias := time.Since(reserva.FechaReserva).Hours() / 24

	// Verifica si han pasado más de 3 días
	if diferenciaDias > 3 {
		http.Error(w, "No se puede borrar la reserva, han pasado más de 3 días desde la reserva", http.StatusForbidden)
		return
	}

	// Borra la reserva si la diferencia de días es 3 o menos
	db.Delete(&reserva)

	// Respuesta exitosa con un mensaje que incluye el ID borrado
	mensaje := fmt.Sprintf("Reserva con ID %d borrada exitosamente", idInt)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mensaje": mensaje})
}
