package service

import (
	"context"
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

	requestedDateDayOfWeek := r.processDayOfWeek(requestedDate)
	requestedDateWeekOfMonth := r.processWeekOfMonth(requestedDate)

	for _, regulation := range segment.Regulations {
		for _, slot := range regulation.RegulatedSlots {
			startHour, startMinute, _ := slot.StartTime.Clock()
			startTime := time.Date(requestedDate.Year(), requestedDate.Month(), requestedDate.Day(), startHour, startMinute, 0, 0, requestedDate.Location())

			endHour, endMinute, _ := slot.EndTime.Clock()
			endTime := time.Date(requestedDate.Year(), requestedDate.Month(), requestedDate.Day(), endHour, endMinute, 0, 0, requestedDate.Location())

			if requestedDate.After(startTime) && requestedDate.Before(endTime) { // If the requested time is inside a slot
				ar := model.ActiveRegulation{
					ID:               regulation.ID,
					Name:             regulation.Name,
					StartTime:        startTime,
					EndTime:          endTime,
					AllowedThreshold: slot.AllowedThreshold,
					Cost:             slot.Cost,
				}
				if slot.IsDaily { // If daily
					activeRegulations = append(activeRegulations, ar)
				} else if slot.DayOfWeek != nil && *slot.DayOfWeek == requestedDateDayOfWeek { // If it has a day of week
					if slot.WeekOfMonth != nil { // If it has a week of month
						if *slot.WeekOfMonth == requestedDateWeekOfMonth {
							activeRegulations = append(activeRegulations, ar)
						}
					} else {
						activeRegulations = append(activeRegulations, ar)
					}
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
	beginningOfTheMonth := time.Date(current.Year(), current.Month(), 1, 0, 0, 0, 0, current.Location())
	_, thisWeek := current.ISOWeek()
	_, beginningWeek := beginningOfTheMonth.ISOWeek()
	return 1 + thisWeek - beginningWeek
}
