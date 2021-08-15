package controller

import (
	"net/http"

	"github.com/Joker666/spotangels-demo/web/resp"
)

// RegulationController holds necessary fields and data for serving regulation handlers
type RegulationController struct {
}

// NewRegulationController returns a new instance of RegulationController
func NewRegulationController() *RegulationController {
	return &RegulationController{}
}

// GetActiveRegulation returns active regulation for a time and spot
func (rc *RegulationController) GetActiveRegulation(w http.ResponseWriter, r *http.Request) {
	resp.ServeData(w, r, http.StatusOK, nil)
	return
}
