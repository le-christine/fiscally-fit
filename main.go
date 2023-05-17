package main

import (
	"github.com/le-christine/fiscally-fit/app/handlers"
	"github.com/le-christine/fiscally-fit/app/models"
	"github.com/le-christine/fiscally-fit/app/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("Tax Calculator Application Starting")

	taxService := services.NewTaxService(models.AllTaxBrackets)
	taxHandler := handlers.NewTaxHandler(taxService)

	api := mux.NewRouter()
	api.HandleFunc("/health", HealthCheckHandler).Methods(http.MethodGet)
	api.HandleFunc("/calculate", taxHandler.CalculateTax).Queries("income", "{income}", "year", "{year}").Methods(http.MethodGet)

	log.Println("Application starting on 8080")
	log.Fatal(http.ListenAndServe(":8080",api))

}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for Health Check")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(`{"alive": true}`))

	if err != nil {
		log.Println("There was an error writing for Health Check", err)
	}

}