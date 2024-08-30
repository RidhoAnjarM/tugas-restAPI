package controllers

import (
	"encoding/json"
	"main/models"
	"net/http"
)

func (idb InDb) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	user_list := []models.User{}

	idb.sql.Find(&user_list)

	if len(user_list) == 0 {
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
		"data":    user_list,
	})
}


