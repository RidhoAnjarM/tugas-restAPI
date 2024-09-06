package models

import (
	"fmt"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetSqlConnection() (*gorm.DB, error) {
	// Membuat DSN (Data Source Name) untuk PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASS"),
		os.Getenv("PG_DB"),
		os.Getenv("PG_PORT"),
	)

	// Membuka koneksi ke PostgreSQL dengan GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Log jika koneksi berhasil
	fmt.Println("Berhasil terhubung ke PostgreSQL!")

	return db, nil
}
