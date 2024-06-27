package services

import (
	"errors"
	"exoplanet/internal/models"
)

func (s *exoplanetService) AddExoplanet(exoplanet models.Exoplanet) (models.Exoplanet, error) {
	if err := validateExoplanet(exoplanet); err != nil {
		return exoplanet, err
	}
	return s.repo.Add(exoplanet)
}

func (s *exoplanetService) ListExoplanets() []models.Exoplanet {
	return s.repo.GetAll()
}

func (s *exoplanetService) GetExoplanet(id string) (models.Exoplanet, error) {
	return s.repo.GetByID(id)
}

func (s *exoplanetService) UpdateExoplanet(id string, exoplanet models.Exoplanet) (models.Exoplanet, error) {
	if err := validateExoplanet(exoplanet); err != nil {
		return exoplanet, err
	}
	return s.repo.Update(id, exoplanet)
}

func (s *exoplanetService) DeleteExoplanet(id string) error {
	return s.repo.Delete(id)
}

// Helper functions
func validateExoplanet(exoplanet models.Exoplanet) error {
	if exoplanet.Name == "" || exoplanet.Description == "" || exoplanet.Distance < 10 || exoplanet.Distance > 1000 || exoplanet.Radius < 0.1 || exoplanet.Radius > 10 {
		return errors.New("invalid exoplanet data")
	}
	if exoplanet.Type == "Terrestrial" && (exoplanet.Mass < 0.1 || exoplanet.Mass > 10) {
		return errors.New("invalid mass for terrestrial planet")
	}
	if exoplanet.Type != "GasGiant" && exoplanet.Type != "Terrestrial" {
		return errors.New("invalid exoplanet type")
	}
	return nil
}
