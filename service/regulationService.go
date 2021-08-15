package service

import (
	"context"
	"log"
	"time"

	"github.com/Joker666/spotangels-demo/model"
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

func (r *RegulationService) GetRegulatedSlotsForATime(ctx context.Context, segmentID int, requestedDate time.Time) (*model.ActiveRegulation, []model.ActiveRegulation, error) {
	var activeRegulations []model.ActiveRegulation
	var highPriorityActiveRegulation *model.ActiveRegulation

	segment, err := r.r.GetBySegment(ctx, segmentID)
	if err != nil {
		return highPriorityActiveRegulation, activeRegulations, err
	}

	now := time.Now()
	requestedDateDayOfWeek := r.processDayOfWeek(requestedDate)
	log.Println(requestedDateDayOfWeek)

	for _, regulation := range segment.Regulations {
		for _, slot := range regulation.RegulatedSlots {
			startHour, startMinute, _ := slot.StartTime.Clock()
			startTime := time.Date(now.Year(), now.Month(), now.Day(), startHour, startMinute, 0, 0, now.Location())

			endHour, endMinute, _ := slot.EndTime.Clock()
			endTime := time.Date(now.Year(), now.Month(), now.Day(), endHour, endMinute, 0, 0, now.Location())

			if requestedDate.After(startTime) && requestedDate.Before(endTime) {
				ar := model.ActiveRegulation{
					ID:               regulation.ID,
					Name:             regulation.Name,
					StartTime:        startTime,
					EndTime:          endTime,
					AllowedThreshold: slot.AllowedThreshold,
					Cost:             slot.Cost,
				}
				if slot.IsDaily {
					activeRegulations = append(activeRegulations, ar)
				} else if slot.DayOfWeek != nil && *slot.DayOfWeek == requestedDateDayOfWeek {
					activeRegulations = append(activeRegulations, ar)
				}
			}
		}
	}

	for i, activeRegulation := range activeRegulations {
		if activeRegulation.AllowedThreshold == nil {
			highPriorityActiveRegulation = &activeRegulations[i]
			break
		}
	}

	return highPriorityActiveRegulation, activeRegulations, nil
}

func (r *RegulationService) processDayOfWeek(current time.Time) int {
	if current.Weekday() == time.Sunday {
		return 7
	}
	return int(current.Weekday())
}

func (r *RegulationService) processWeekOfMonth(current time.Time) int {
	if current.Weekday() == time.Sunday {
		return 7
	}
	return int(current.Weekday())
}
