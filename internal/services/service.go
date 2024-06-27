package services

import (
	"exoplanet/internal/models"
	"exoplanet/internal/repositories"
)

type ExoplanetService interface {
	AddExoplanet(exoplanet models.Exoplanet) (models.Exoplanet, error)
	ListExoplanets() []models.Exoplanet
	GetExoplanet(id string) (models.Exoplanet, error)
	UpdateExoplanet(id string, exoplanet models.Exoplanet) (models.Exoplanet, error)
	DeleteExoplanet(id string) error
	CalculateFuelCost(id string, crewCapacity int) (float64, error)
}

type exoplanetService struct {
	repo repositories.ExoplanetRepository
}

func NewExoplanetService(repo repositories.ExoplanetRepository) ExoplanetService {
	return &exoplanetService{repo: repo}
}
