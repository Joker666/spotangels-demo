package resp

import "github.com/Joker666/spotangels-demo/model"

// User holds necessary fields for a user
type ActiveRegulations struct {
	HighPriority *model.ActiveRegulation  `json:"high_priority"`
	All          []model.ActiveRegulation `json:"all"`
}
