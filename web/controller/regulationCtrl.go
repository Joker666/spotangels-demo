package controller

import (
	"log"
	"net/http"
	"strconv"
	"time"

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
	segmentID, err := strconv.Atoi(r.URL.Query().Get("segment_id"))

	dtl := map[string]interface{}{}
	if err != nil || segmentID < 1 {
		dtl["segment_id"] = "is invalid"
		resp.ServeUnprocessableEntity(w, r, ErrInvalidData, dtl)
		return
	}

	timestamp, err := strconv.Atoi(r.URL.Query().Get("timestamp"))
	if err != nil || timestamp < 1 {
		dtl["timestamp"] = "is invalid"
		resp.ServeUnprocessableEntity(w, r, ErrInvalidData, dtl)
		return
	}

	requestedDate := time.Unix(int64(timestamp), 0)
	log.Println("Requested Date: ", requestedDate)

	highPriorityActiveRegulation, activeRegulations, err := rc.s.GetRegulatedSlotsForATime(ctx, segmentID, requestedDate)
	if err != nil {
		resp.ServeInternalServerError(w, r, ErrProcessFailure)
		return
	}

	data := resp.ActiveRegulations{HighPriority: highPriorityActiveRegulation, All: activeRegulations}
	resp.ServeData(w, r, http.StatusOK, data)
}
