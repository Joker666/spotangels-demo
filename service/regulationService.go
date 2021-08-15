package service

import (
	"github.com/Joker666/spotangels-demo/repo"
)

// RegulationService is the service that handles showing the regulated times for parking slots
type RegulationService struct {
	r repo.RegulationRepository
}

// NewRegulationService returns a new instance of regulation service
func NewRegulationService(r repo.RegulationRepository) *RegulationService {
	return &RegulationService{
		r: r,
	}
}
