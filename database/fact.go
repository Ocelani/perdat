package database

import "time"

// Fact is a lifetime register of an event.
type Fact struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	DateTime  time.Time // (default time.Now)
	CratedAt  time.Time
	UpdatedAt time.Time
}

// NewFact instantiates a new fact.
func NewFact(name string, date time.Time) *Fact {
	return &Fact{
		Name:     name,
		DateTime: date,
	}
}
