package repositories

import (
	"exoplanet/internal/models"
)

type ExoplanetRepository interface {
	Add(exoplanet models.Exoplanet) (models.Exoplanet, error)
	GetAll() []models.Exoplanet
	GetByID(id string) (models.Exoplanet, error)
	Update(id string, exoplanet models.Exoplanet) (models.Exoplanet, error)
	Delete(id string) error
}

type InMemoryExoplanetRepository struct {
	exoplanets map[string]models.Exoplanet
}

func NewInMemoryExoplanetRepository() ExoplanetRepository {
	return &InMemoryExoplanetRepository{
		exoplanets: make(map[string]models.Exoplanet),
	}
}
