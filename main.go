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

	// Initialize repository, service, and controller
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	//To be able to handle different requests -- create router -- mux
	router := mux.NewRouter()
	router.HandleFunc("/users", userController.GetAllUsers).Methods("GET")

	//router.HandleFunc("/users", getUsers(db)).Methods("GET")
	//router.HandleFunc("users/{id}", getUser(db)).Methods("GET")
	//router.HandleFunc("/users", createUser(db)).Methods("POST")
	//router.HandleFunc("/users/{id}", updateUser(db)).Methods("PUT")
	//router.HandleFunc("/users/{id}", deleteUser(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
