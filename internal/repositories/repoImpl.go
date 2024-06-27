package repositories

import (
	"errors"
	"exoplanet/internal/models"
	"strconv"
)

func (repo *InMemoryExoplanetRepository) Add(exoplanet models.Exoplanet) (models.Exoplanet, error) {
	exoplanet.ID = strconv.Itoa(len(repo.exoplanets) + 1)
	repo.exoplanets[exoplanet.ID] = exoplanet
	return exoplanet, nil
}

func (repo *InMemoryExoplanetRepository) GetAll() []models.Exoplanet {
	var list []models.Exoplanet
	for _, exoplanet := range repo.exoplanets {
		list = append(list, exoplanet)
	}
	return list
}

func (repo *InMemoryExoplanetRepository) GetByID(id string) (models.Exoplanet, error) {
	exoplanet, exists := repo.exoplanets[id]
	if !exists {
		return exoplanet, errors.New("Exoplanet not found")
	}
	return exoplanet, nil
}

func (repo *InMemoryExoplanetRepository) Update(id string, exoplanet models.Exoplanet) (models.Exoplanet, error) {
	_, exists := repo.exoplanets[id]
	if !exists {
		return exoplanet, errors.New("Exoplanet not found")
	}
	exoplanet.ID = id
	repo.exoplanets[id] = exoplanet
	return exoplanet, nil
}

func (repo *InMemoryExoplanetRepository) Delete(id string) error {
	_, exists := repo.exoplanets[id]
	if !exists {
		return errors.New("Exoplanet not found")
	}
	delete(repo.exoplanets, id)
	return nil
}
