package model

type Regulation struct {
	ID             int             `gorm:"column:id;primary_key" json:"id"`
	Name           string          `gorm:"column:name" json:"email"`
	SegmentID      int             `gorm:"column:segment_id" json:"segment_id"`
	RegulatedSlots []RegulatedSlot `gorm:"foreignKey:RegulationID" json:"regulated_slots"`
}

// TableName sets the insert table name for this struct type
func (a *Regulation) TableName() string {
	return "regulations"
}
