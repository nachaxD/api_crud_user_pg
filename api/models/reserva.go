package models

import "time"

// Reserva representa el modelo de la tabla de reservas
type Reserva struct {
	ID                   uint      `gorm:"primaryKey"`
	IDFechaPaquete       uint      `gorm:"column:id_fechapaquete"`
	PrecioTotal          float64   `gorm:"column:precio_total"`
	FechaReserva         time.Time `gorm:"column:fecha_reserva"`
	ServiciosAdicionales string    `gorm:"column:servicios_adicionales"`
	IDUsuario            string    `gorm:"column:id_usuario"`
	Estado               string    `gorm:"column:estado"`
	Pasajeros            string    `gorm:"column:pasajeros"`
}

// TableName especifica el nombre de la tabla en la base de datos
func (Reserva) TableName() string {
	return "reserva" // Reemplaza "nombre_de_tu_tabla" con el nombre real de tu tabla
}
