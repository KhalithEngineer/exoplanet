package main

import (
	"log"
	"net/http"

	"exoplanet/internal/handlers"
	"exoplanet/internal/repositories"
	"exoplanet/internal/services"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	repo := repositories.NewInMemoryExoplanetRepository()
	service := services.NewExoplanetService(repo)
	handler := handlers.NewHandler(service)

	r.HandleFunc("/exoplanets", handler.AddExoplanet).Methods("POST")
	r.HandleFunc("/exoplanets", handler.ListExoplanets).Methods("GET")
	r.HandleFunc("/exoplanets/{id}", handler.GetExoplanet).Methods("GET")
	r.HandleFunc("/exoplanets/{id}", handler.UpdateExoplanet).Methods("PUT")
	r.HandleFunc("/exoplanets/{id}", handler.DeleteExoplanet).Methods("DELETE")
	r.HandleFunc("/exoplanets/{id}/fuel", handler.FuelEstimation).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
