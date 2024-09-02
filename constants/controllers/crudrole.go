package controllers

import (
	"encoding/json"
	"log"
	"main/models"
	"net/http"

	"gorm.io/gorm"
)


type CRUDROLE struct {
	DB *gorm.DB
}

func validateTokenCrudrole(r *http.Request) bool {
    token := r.Header.Get("Authorization")
    return token == "Bearer token_app"

}

// Fungsi untuk membuat instance CRUDROLE baru
func NewCRUDROLE(db *gorm.DB) *CRUDROLE {
	return &CRUDROLE{DB: db}
}

// Fungsi untuk menambahkan user baru
func (ctrl *CRUDROLE) CreateRole(w http.ResponseWriter, r *http.Request) {
	// Validasi token
	if !validateTokenCrudrole(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Dekode body JSON ke dalam struct User
	var role models.Role
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validasi data user (tambahkan validasi lebih lanjut jika perlu)

	// Simpan user ke dalam database
	if err := ctrl.DB.Create(&role).Error; err != nil {
		log.Printf("Error creating role: %v", err)
		http.Error(w, "Error creating role", http.StatusInternalServerError)
		return
	}

	// Set header dan kirim response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(role)
}

// Fungsi untuk mengambil daftar semua user
func (ctrl *CRUDROLE) ListRole(w http.ResponseWriter, r *http.Request) {
	// Validasi token
	if !validateTokenCrudrole(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Ambil semua user dari database
	var role []models.Role
	if err := ctrl.DB.Find(&role).Error; err != nil {
		log.Printf("Error fetching role: %v", err)
		http.Error(w, "Error fetching role", http.StatusInternalServerError)
		return
	}

	// Set header dan kirim response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(role)
}

// Fungsi untuk mengambil user berdasarkan ID
func (ctrl *CRUDROLE) GetRole(w http.ResponseWriter, r *http.Request) {
	// Validasi token
	if !validateTokenCrudrole(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Ambil ID dari parameter query
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	// Cari role berdasarkan ID
	var role models.Role
	if err := ctrl.DB.First(&role, id).Error; err != nil {
		log.Printf("Error fetching role: %v", err)
		http.Error(w, "role not found", http.StatusNotFound)
		return
	}

	// Set header dan kirim response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(role)
}

// Fungsi untuk memperbarui user berdasarkan ID
func (ctrl *CRUDROLE) UpdateRole(w http.ResponseWriter, r *http.Request) {
	// Validasi token
	if !validateTokenCrudrole(r) {
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
	var role models.Role
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Perbarui user berdasarkan ID
	if err := ctrl.DB.Model(&models.Role{}).Where("id = ?", id).Updates(role).Error; err != nil {
		log.Printf("Error updating role: %v", err)
		http.Error(w, "Error updating role", http.StatusInternalServerError)
		return
	}

	// Set header dan kirim response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(role)
}

// Fungsi untuk menghapus user berdasarkan ID
func (ctrl *CRUDROLE) DeleteRole(w http.ResponseWriter, r *http.Request) {
	// Validasi token
	if !validateTokenCrudrole(r) {
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
	if err := ctrl.DB.Delete(&models.Role{}, id).Error; err != nil {
		log.Printf("Error deleting user: %v", err)
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	// Kembalikan status No Content (204)
	w.WriteHeader(http.StatusNoContent)
}
