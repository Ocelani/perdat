package entity

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Counter is a register of a life event.
type Counter struct {
	gorm.Model
	Name string
}

// NewCounter instantiates a counter object.
func NewCounter(counterName string) *Counter {
	return &Counter{
		Name: counterName,
	}
}

// NewCounters instantiates many counters objects.
func NewCounters(counterNames []string) *[]Counter {
	counters := []Counter{}
	for _, name := range counterNames {
		counters = append(counters, *NewCounter(name))
	}
	return &counters
}

func (f *Counter) String() string {
	return fmt.Sprintf(`
	Name:     %v
	DateTime: %v
	`, f.Name, f.CreatedAt.Format(time.RFC3339))
}
