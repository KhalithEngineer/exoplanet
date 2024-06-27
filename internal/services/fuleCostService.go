package services

import "exoplanet/internal/models"

func (s *exoplanetService) CalculateFuelCost(id string, crewCapacity int) (float64, error) {
	exoplanet, err := s.repo.GetByID(id)
	if err != nil {
		return 0, err
	}
	return calculateFuelCost(exoplanet, crewCapacity), nil
}

func calculateFuelCost(exoplanet models.Exoplanet, crewCapacity int) float64 {
	var gravity float64
	if exoplanet.Type == "GasGiant" {
		gravity = 0.5 / (exoplanet.Radius * exoplanet.Radius)
	} else if exoplanet.Type == "Terrestrial" {
		gravity = exoplanet.Mass / (exoplanet.Radius * exoplanet.Radius)
	}
	fuelCost := float64(exoplanet.Distance) / (gravity * gravity) * float64(crewCapacity)
	return fuelCost
}
