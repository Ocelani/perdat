package pkg

import (
	"fmt"
	"time"

	"github.com/pkg/errors"

	"github.com/Ocelani/perdat/database"
)

// Fact is a lifetime register of an event.
type Fact struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	DateTime  time.Time // (default time.Now)
	CratedAt  time.Time
	UpdatedAt time.Time
}

// NewFact instantiates a fact object.
func NewFact(name string, date time.Time) *Fact {
	return &Fact{
		Name:     name,
		DateTime: date,
	}
}

// Create a completely new fact in database.
func (f *Fact) Create() error {
	db, err := database.Connect()
	if err != nil {
		return errors.Wrap(err, "Couldn't register fact due to")
	}
	db.Create(f)

	return nil
}

func (f *Fact) String() string {
	return fmt.Sprintf(`
	Name:     %v
	DateTime: %v
	`, f.Name, f.DateTime)
}
