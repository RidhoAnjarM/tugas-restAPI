package controllers

import (
	"encoding/json"
	"log"
	"main/models"
	"net/http"

	"gorm.io/gorm"
)


type CRUDAC struct {
	DB *gorm.DB
}

func validateTokenCrudAC(r *http.Request) bool {
    token := r.Header.Get("Authorization")
    return token == "Bearer token_app"

}

// Fungsi untuk membuat instance CRUDAC baru
func NewCRUDAC(db *gorm.DB) *CRUDAC {
	return &CRUDAC{DB: db}
}

// Fungsi untuk menambahkan user baru
func (ctrl *CRUDAC) CreateAC(w http.ResponseWriter, r *http.Request) {
	// Validasi token
	if !validateTokenCrudAC(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Dekode body JSON ke dalam struct User
	var ac models.AC
	if err := json.NewDecoder(r.Body).Decode(&ac); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validasi data user (tambahkan validasi lebih lanjut jika perlu)

	// Simpan user ke dalam database
	if err := ctrl.DB.Create(&ac).Error; err != nil {
		log.Printf("Error creating ac: %v", err)
		http.Error(w, "Error creating ac", http.StatusInternalServerError)
		return
	}

	// Set header dan kirim response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ac)
}

// Fungsi untuk mengambil daftar semua user
func (ctrl *CRUDAC) ListAC(w http.ResponseWriter, r *http.Request) {
	// Validasi token
	if !validateTokenCrudAC(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Ambil semua user dari database
	var ac []models.AC
	if err := ctrl.DB.Find(&ac).Error; err != nil {
		log.Printf("Error fetching ac: %v", err)
		http.Error(w, "Error fetching ac", http.StatusInternalServerError)
		return
	}

	// Set header dan kirim response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ac)
}

// Fungsi untuk mengambil user berdasarkan ID
func (ctrl *CRUDAC) GetAC(w http.ResponseWriter, r *http.Request) {
	// Validasi token
	if !validateTokenCrudAC(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Ambil ID dari parameter query
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	// Cari ac berdasarkan ID
	var ac models.AC
	if err := ctrl.DB.First(&ac, id).Error; err != nil {
		log.Printf("Error fetching ac: %v", err)
		http.Error(w, "ac not found", http.StatusNotFound)
		return
	}

	// Set header dan kirim response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ac)
}

// Fungsi untuk memperbarui user berdasarkan ID
func (ctrl *CRUDAC) UpdateAC(w http.ResponseWriter, r *http.Request) {
	// Validasi token
	if !validateTokenCrudAC(r) {
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
	var ac models.AC
	if err := json.NewDecoder(r.Body).Decode(&ac); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Perbarui user berdasarkan ID
	if err := ctrl.DB.Model(&models.AC{}).Where("id = ?", id).Updates(ac).Error; err != nil {
		log.Printf("Error updating ac: %v", err)
		http.Error(w, "Error updating ac", http.StatusInternalServerError)
		return
	}

	// Set header dan kirim response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ac)
}

// Fungsi untuk menghapus user berdasarkan ID
func (ctrl *CRUDAC) DeleteAC(w http.ResponseWriter, r *http.Request) {
	// Validasi token
	if !validateTokenCrudAC(r) {
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
	if err := ctrl.DB.Delete(&models.AC{}, id).Error; err != nil {
		log.Printf("Error deleting user: %v", err)
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	// Kembalikan status No Content (204)
	w.WriteHeader(http.StatusNoContent)
}
