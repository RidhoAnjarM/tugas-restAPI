package controllers

import (
	"encoding/json"
	"main/models"
	"net/http"
)

func (idb InDb) GetAllService(w http.ResponseWriter, r *http.Request) {
	service_list := []models.Service{}

	idb.sql.Find(&service_list)

	if len(service_list) == 0 {
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
		"data":    service_list,
	})
}