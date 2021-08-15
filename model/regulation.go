package model

type Regulation struct {
	ID        int    `gorm:"column:id;primary_key" json:"id"`
	Name      string `gorm:"column:name" json:"email"`
	SegmentID int    `gorm:"column:segment_id" json:"segment_id"`
}

type RegulationInfo struct {
	Segment                Segment        `gorm:"embedded"`
	Regulation             Regulation     `gorm:"embedded"`
	PlatRegulationSlotform RegulationSlot `gorm:"embedded"`
}
