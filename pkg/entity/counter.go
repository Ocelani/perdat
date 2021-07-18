package entity

import (
	"gorm.io/gorm"
)

// Counter is a register of a life event.
type Counter struct {
	gorm.Model
	Name     string
	CountNum int
}

// NewCounter instantiates a counter object.
func NewCounter(counterName string) *Counter {
	return &Counter{
		Name: counterName,
	}
}

// NewCounters instantiates many counters objects.
func NewCounters(counterNames []string, countNum int) *[]Counter {
	counters := []Counter{}
	for _, name := range counterNames {
		counters = append(counters, *NewCounter(name))
	}
	return &counters
}
