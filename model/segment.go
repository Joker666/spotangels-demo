package model

type Segment struct {
	ID          int          `gorm:"column:id;primary_key" json:"id"`
	Name        string       `gorm:"column:name" json:"email"`
	Regulations []Regulation `gorm:"foreignKey:SegmentID" json:"regulations"`
}

// TableName sets the insert table name for this struct type
func (a *Segment) TableName() string {
	return "segments"
}
