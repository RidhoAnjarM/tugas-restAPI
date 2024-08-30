package controllers

import (
	"encoding/json"
	"net/http"
	"main/models"
	"gorm.io/gorm"
	"log"
)

// Fungsi untuk validasi token
func validateTokenCrud(r *http.Request) bool {
    token := r.Header.Get("Authorization")
    return token == "Bearer token_app"
}

type CRUDController struct {
	DB *gorm.DB
}

// Fungsi untuk membuat instance CRUDController baru
func NewCRUDController(db *gorm.DB) *CRUDController {
	return &CRUDController{DB: db}
}

// Fungsi untuk menambahkan user baru
func (ctrl *CRUDController) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Validasi token
	if !validateTokenCrud(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Dekode body JSON ke dalam struct User
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validasi data user (tambahkan validasi lebih lanjut jika perlu)

	// Simpan user ke dalam database
	if err := ctrl.DB.Create(&user).Error; err != nil {
		log.Printf("Error creating user: %v", err)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// Set header dan kirim response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Fungsi untuk mengambil daftar semua user
func (ctrl *CRUDController) ListUsers(w http.ResponseWriter, r *http.Request) {
	// Validasi token
	if !validateTokenCrud(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Ambil semua user dari database
	var users []models.User
	if err := ctrl.DB.Find(&users).Error; err != nil {
		log.Printf("Error fetching users: %v", err)
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}

	// Set header dan kirim response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Fungsi untuk mengambil user berdasarkan ID
func (ctrl *CRUDController) GetUser(w http.ResponseWriter, r *http.Request) {
	// Validasi token
	if !validateTokenCrud(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Ambil ID dari parameter query
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	// Cari user berdasarkan ID
	var user models.User
	if err := ctrl.DB.First(&user, id).Error; err != nil {
		log.Printf("Error fetching user: %v", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Set header dan kirim response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Fungsi untuk memperbarui user berdasarkan ID
func (ctrl *CRUDController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Validasi token
	if !validateTokenCrud(r) {
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
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Perbarui user berdasarkan ID
	if err := ctrl.DB.Model(&models.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		log.Printf("Error updating user: %v", err)
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	// Set header dan kirim response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Fungsi untuk menghapus user berdasarkan ID
func (ctrl *CRUDController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Validasi token
	if !validateTokenCrud(r) {
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
	if err := ctrl.DB.Delete(&models.User{}, id).Error; err != nil {
		log.Printf("Error deleting user: %v", err)
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	// Kembalikan status No Content (204)
	w.WriteHeader(http.StatusNoContent)
}
