package controller

import (
	"net/http"

	"github.com/Joker666/spotangels-demo/service"
	"github.com/Joker666/spotangels-demo/web/resp"
)

// RegulationController holds necessary fields and data for serving regulation handlers
type RegulationController struct {
	s *service.RegulationService
}

// NewRegulationController returns a new instance of RegulationController
func NewRegulationController(regulationService *service.RegulationService) *RegulationController {
	return &RegulationController{
		s: regulationService,
	}
}

// GetActiveRegulation returns active regulation for a time and spot
func (rc *RegulationController) GetActiveRegulation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rc.s.GetRegulatedSlotsForATime(ctx, 4)
	resp.ServeData(w, r, http.StatusOK, nil)
	return
}
