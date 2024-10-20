package main

import (
	"api/config"
	"api/controller"
	"api/repository"
	"api/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// Connect to the database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY,name TEXT, email TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize repository, service, and controller
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	//To be able to handle different requests -- create router -- mux
	router := mux.NewRouter()
	router.HandleFunc("/users", userController.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userController.GetUserByID).Methods("GET")
	router.HandleFunc("/users", userController.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userController.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
