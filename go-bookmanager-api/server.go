package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"matheusmv.com/go-bookmanager-api/config"
	"matheusmv.com/go-bookmanager-api/database"
	"matheusmv.com/go-bookmanager-api/models"
	"matheusmv.com/go-bookmanager-api/repositories"
	"matheusmv.com/go-bookmanager-api/resources"
	"matheusmv.com/go-bookmanager-api/services"
)

func main() {

	appConfig, err := config.LoadConfigurations()

	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect(appConfig)

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Author{}, &models.Book{})

	bookRepository := repositories.NewBookRepository(db)
	bookService := services.NewService(bookRepository)
	bookResources := resources.NewBookResourceHandler(bookService)

	routes := mux.NewRouter()

	routes.HandleFunc("/books", bookResources.GetBooks).Methods("GET")
	routes.HandleFunc("/books/{id}", bookResources.GetBook).Methods("GET")
	routes.HandleFunc("/books", bookResources.CreateBook).Methods("POST")
	routes.HandleFunc("/books/{id}", bookResources.UpdateBook).Methods("PUT")
	routes.HandleFunc("/books/{id}", bookResources.DeleteBook).Methods("DELETE")

	port := fmt.Sprintf(":%d", appConfig.Server.Port)

	log.Fatal(http.ListenAndServe(port, routes))
}
