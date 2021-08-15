package service

import (
	"context"
	"log"

	"github.com/Joker666/spotangels-demo/repo"
	"github.com/davecgh/go-spew/spew"
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

func (r *RegulationService) GetRegulatedSlotsForATime(ctx context.Context, segmentID int) {
	info, err := r.r.GetBySegment(ctx, segmentID)
	spew.Dump(info)
	spew.Dump(err)
	log.Println("Okey")
}
