package handlers

import (
	"exoplanet/internal/services"
)

type Handler struct {
	service services.ExoplanetService
}

func NewHandler(service services.ExoplanetService) *Handler {
	return &Handler{service: service}
}
