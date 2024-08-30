package main

import (
	"fmt"
	"net"
	"net/http"
	"main/controllers"
	"main/models"
	"main/middleware"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(err.Error())
		return
	}

	db, err := models.GetSqlConnection()
	if err != nil {
		fmt.Println("gagal koneksi database:", err)
		return
	}

	userController := controllers.NewCRUDController(db)

	http.HandleFunc("/api/v1/users", middleware.TokenValidation(userController.ListUsers))
	http.HandleFunc("/api/v1/users/get", middleware.TokenValidation(userController.GetUser)) // Endpoint untuk mendapatkan user berdasarkan ID
	http.HandleFunc("/api/v1/users/create", middleware.TokenValidation(userController.CreateUser))
	http.HandleFunc("/api/v1/users/update", middleware.TokenValidation(userController.UpdateUser))
	http.HandleFunc("/api/v1/users/delete", middleware.TokenValidation(userController.DeleteUser))


	ln, err := net.Listen("tcp", "localhost:3333")
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}

	fmt.Println("Server is running on http://localhost:3333")
	if err := http.Serve(ln, nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
