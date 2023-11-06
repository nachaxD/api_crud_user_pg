package utils

import (
	"backend/api/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// OpenDB abre la conexión con la base de datos y la devuelve como una instancia de GORM
func OpenDB() (*gorm.DB, error) {
	dsn := config.DBURL() // Asegúrate de configurar tu DSN según tus necesidades

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
