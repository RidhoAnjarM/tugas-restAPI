package controllers

import (
	"encoding/json"
	"main/models"
	"net/http"
)

func (idb InDb) GetAllAC(w http.ResponseWriter, r *http.Request) {
	ac_list := []models.AC{}

	idb.sql.Find(&ac_list)

	if len(ac_list) == 0 {
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
		"data":    ac_list,
	})
}