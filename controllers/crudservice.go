package controllers

import (
	"encoding/json"
	"log"
	"main/models"
	"net/http"

	"gorm.io/gorm"
)


type CRUDSERVICE struct {
	DB *gorm.DB
}

func validateTokenCrudservice(r *http.Request) bool {
    token := r.Header.Get("Authorization")
    return token == "Bearer token_app"

}

// Fungsi untuk membuat instance CRUDSERVICE baru
func NewCRUDSERVICE(db *gorm.DB) *CRUDSERVICE {
	return &CRUDSERVICE{DB: db}
}

// Fungsi untuk menambahkan user baru
func (ctrl *CRUDSERVICE) CreateService(w http.ResponseWriter, r *http.Request) {
	// Validasi token
	if !validateTokenCrudservice(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Dekode body JSON ke dalam struct User
	var service models.Service
	if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validasi data user (tambahkan validasi lebih lanjut jika perlu)

	// Simpan user ke dalam database
	if err := ctrl.DB.Create(&service).Error; err != nil {
		log.Printf("Error creating service: %v", err)
		http.Error(w, "Error creating service", http.StatusInternalServerError)
		return
	}

	// Set header dan kirim response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(service)
}

// Fungsi untuk mengambil daftar semua user
func (ctrl *CRUDSERVICE) ListService(w http.ResponseWriter, r *http.Request) {
	// Validasi token
	if !validateTokenCrudservice(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Ambil semua user dari database
	var services []models.Service
	if err := ctrl.DB.Find(&services).Error; err != nil {
		log.Printf("Error fetching services: %v", err)
		http.Error(w, "Error fetching services", http.StatusInternalServerError)
		return
	}

	// Set header dan kirim response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

// Fungsi untuk mengambil user berdasarkan ID
func (ctrl *CRUDSERVICE) GetService(w http.ResponseWriter, r *http.Request) {
	// Validasi token
	if !validateTokenCrudservice(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Ambil ID dari parameter query
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	// Cari service berdasarkan ID
	var service models.Service
	if err := ctrl.DB.First(&service, id).Error; err != nil {
		log.Printf("Error fetching service: %v", err)
		http.Error(w, "service not found", http.StatusNotFound)
		return
	}

	// Set header dan kirim response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service)
}

// Fungsi untuk memperbarui user berdasarkan ID
func (ctrl *CRUDSERVICE) UpdateService(w http.ResponseWriter, r *http.Request) {
	// Validasi token
	if !validateTokenCrudservice(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Ambil ID dari parameter query
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	// Dekode body JSON ke dalam struct User
	var service models.Service
	if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Perbarui user berdasarkan ID
	if err := ctrl.DB.Model(&models.Service{}).Where("id = ?", id).Updates(service).Error; err != nil {
		log.Printf("Error updating service: %v", err)
		http.Error(w, "Error updating service", http.StatusInternalServerError)
		return
	}

	// Set header dan kirim response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service)
}

// Fungsi untuk menghapus user berdasarkan ID
func (ctrl *CRUDSERVICE) DeleteService(w http.ResponseWriter, r *http.Request) {
	// Validasi token
	if !validateTokenCrudservice(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Ambil ID dari parameter query
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	// Hapus user berdasarkan ID
	if err := ctrl.DB.Delete(&models.Service{}, id).Error; err != nil {
		log.Printf("Error deleting user: %v", err)
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	// Kembalikan status No Content (204)
	w.WriteHeader(http.StatusNoContent)
}
