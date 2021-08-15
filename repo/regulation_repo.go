package repo

import (
	"context"

	"github.com/Joker666/spotangels-demo/model"
	"gorm.io/gorm"
)

// RegulationRepository interface represents the regulation repository
type RegulationRepository interface {
	GetBySegment(ctx context.Context, segmentID int) (*model.Segment, error)
}

// Regulation holds the necessary fields for RegulationRepository
type Regulation struct {
	db *gorm.DB
}

// NewRegulation returns a new instance of Regulation
func NewRegulation(db *gorm.DB) *Regulation {
	return &Regulation{
		db: db,
	}
}

// GetBySegment gets regulationInfo by segmentID
func (u *Regulation) GetBySegment(ctx context.Context, segmentID int) (*model.Segment, error) {
	tx := u.db.WithContext(ctx)

	var segment model.Segment
	err := tx.Preload("Regulations").Preload("Regulations.RegulatedSlots").First(&segment, segmentID).Error
	return &segment, err
}
