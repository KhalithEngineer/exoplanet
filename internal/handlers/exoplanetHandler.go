package handlers

import (
	"encoding/json"
	"exoplanet/internal/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) AddExoplanet(w http.ResponseWriter, r *http.Request) {
	var exoplanet models.Exoplanet
	if err := json.NewDecoder(r.Body).Decode(&exoplanet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	exoplanet, err := h.service.AddExoplanet(exoplanet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(exoplanet)
}

func (h *Handler) ListExoplanets(w http.ResponseWriter, r *http.Request) {
	exoplanets := h.service.ListExoplanets()
	json.NewEncoder(w).Encode(exoplanets)
}

func (h *Handler) GetExoplanet(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	exoplanet, err := h.service.GetExoplanet(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(exoplanet)
}

func (h *Handler) UpdateExoplanet(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var exoplanet models.Exoplanet
	if err := json.NewDecoder(r.Body).Decode(&exoplanet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	exoplanet, err := h.service.UpdateExoplanet(id, exoplanet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(exoplanet)
}

func (h *Handler) DeleteExoplanet(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := h.service.DeleteExoplanet(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) FuelEstimation(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	crewCapacity, err := strconv.Atoi(r.URL.Query().Get("crewCapacity"))
	if err != nil {
		http.Error(w, "Invalid crew capacity", http.StatusBadRequest)
		return
	}
	fuelCost, err := h.service.CalculateFuelCost(id, crewCapacity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"fuelCost": fuelCost})
}
