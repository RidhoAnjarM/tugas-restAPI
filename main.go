package main

import (
	"fmt"
	"net"
	"net/http"
	"main/controllers"
	"main/models"
	"main/middleware"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(err.Error())
		return
	}

	// Membuka koneksi database
	db, err := models.GetSqlConnection()
	if err != nil {
		log.Fatal("Could not connect to the database", err)
	}

	db.AutoMigrate(&models.User{}, &models.Role{}, &models.AC{}, &models.Service{})


	userController := controllers.NewCRUDController(db)

	http.HandleFunc("/api/v1/users", middleware.TokenValidation(userController.ListUsers))
	http.HandleFunc("/api/v1/users/get", middleware.TokenValidation(userController.GetUser))
	http.HandleFunc("/api/v1/users/create", middleware.TokenValidation(userController.CreateUser))
	http.HandleFunc("/api/v1/users/update", middleware.TokenValidation(userController.UpdateUser))
	http.HandleFunc("/api/v1/users/delete", middleware.TokenValidation(userController.DeleteUser))


	roleController := controllers.NewCRUDROLE(db)

	http.HandleFunc("/api/v1/roles", middleware.TokenValidation(roleController.ListRole))
	http.HandleFunc("/api/v1/roles/get", middleware.TokenValidation(roleController.GetRole))
	http.HandleFunc("/api/v1/roles/create", middleware.TokenValidation(roleController.CreateRole))
	http.HandleFunc("/api/v1/roles/update", middleware.TokenValidation(roleController.UpdateRole))
	http.HandleFunc("/api/v1/roles/delete", middleware.TokenValidation(roleController.DeleteRole))

	acController := controllers.NewCRUDAC(db)

	http.HandleFunc("/api/v1/acs", middleware.TokenValidation(acController.ListAC))
	http.HandleFunc("/api/v1/acs/get", middleware.TokenValidation(acController.GetAC))
	http.HandleFunc("/api/v1/acs/create", middleware.TokenValidation(acController.CreateAC))
	http.HandleFunc("/api/v1/acs/update", middleware.TokenValidation(acController.UpdateAC))
	http.HandleFunc("/api/v1/acs/delete", middleware.TokenValidation(acController.DeleteAC))

	serviceController := controllers.NewCRUDSERVICE(db)

	http.HandleFunc("/api/v1/services", middleware.TokenValidation(serviceController.ListService))
	http.HandleFunc("/api/v1/services/get", middleware.TokenValidation(serviceController.GetService))
	http.HandleFunc("/api/v1/services/create", middleware.TokenValidation(serviceController.CreateService))
	http.HandleFunc("/api/v1/services/update", middleware.TokenValidation(serviceController.UpdateService))
	http.HandleFunc("/api/v1/services/delete", middleware.TokenValidation(serviceController.DeleteService))


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
