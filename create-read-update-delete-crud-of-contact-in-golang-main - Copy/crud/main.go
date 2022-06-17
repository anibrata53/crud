package main

import (
	"crud/controllers"
	"crud/database"
	"crud/entity"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {

	initDB()
	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreatePaitent).Methods("POST")
	router.HandleFunc("/get", controllers.GetAllPaitent).Methods("GET")
	router.HandleFunc("/get/{id}", controllers.GetPaitentByID).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdatePaitentByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletPaitentByID).Methods("DELETE")
	router.HandleFunc("/login/{id}", controllers.Login).Methods("POST")
}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "Abittu@2000",
			DB:         "contacts",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Paitent{})
	// database.Migratee(&entity.Login{})
}
