package entity

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Fact is a register of a life event.
type Fact struct {
	gorm.Model
	Name string
}

// NewFact instantiates a fact object.
func NewFact(factName string) *Fact {
	return &Fact{
		Name: factName,
	}
}

// NewFacts instantiates many facts objects.
func NewFacts(factNames []string) *[]Fact {
	facts := []Fact{}
	for _, name := range factNames {
		facts = append(facts, *NewFact(name))
	}
	return &facts
}

func (f *Fact) String() string {
	return fmt.Sprintf(`
	Name:     %v
	DateTime: %v
	`, f.Name, f.CreatedAt.Format(time.RFC3339))
}
