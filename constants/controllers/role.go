package controllers

import (
	"encoding/json"
	"main/models"
	"net/http"
)

func (idb InDb) GetAllRole(w http.ResponseWriter, r *http.Request) {
	role_list := []models.Role{}

	idb.sql.Find(&role_list)

	if len(role_list) == 0 {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  false,
			"message": "Tidak ada user yang ditemukan",
			"data":    nil,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Berhasil",
		"data":    role_list,
	})
}