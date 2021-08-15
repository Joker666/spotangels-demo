package model

import "time"

type RegulationSlot struct {
	ID               int       `gorm:"column:id;primary_key" json:"id"`
	DayOfWeek        *int      `gorm:"column:day_of_week" json:"day_of_week"`
	WeekOfMonth      *int      `gorm:"column:week_of_month" json:"week_of_month"`
	DayOfMonth       *int      `gorm:"column:day_of_month" json:"day_of_month"`
	StartTime        time.Time `gorm:"column:start_time" json:"start_time"`
	EndTime          time.Time `gorm:"column:end_time" json:"end_time"`
	AllowedThreshold *int      `gorm:"column:allowed_threshold" json:"allowed_threshold"`
	Cost             float32   `gorm:"column:cost" json:"cost"`
	IsDaily          bool      `gorm:"column:is_daily" json:"is_daily"`
	RegulationId     int       `gorm:"column:regulation_id" json:"regulation_id"`
}
