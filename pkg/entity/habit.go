package entity

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Habit is a register of a life event.
type Habit struct {
	gorm.Model
	Name string
}

// NewHabit instantiates a habit object.
func NewHabit(habitName string) *Habit {
	return &Habit{
		Name: habitName,
	}
}

// NewHabits instantiates many habits objects.
func NewHabits(habitNames []string) *[]Habit {
	habits := []Habit{}
	for _, name := range habitNames {
		habits = append(habits, *NewHabit(name))
	}
	return &habits
}

func (f *Habit) String() string {
	return fmt.Sprintf(`
	Name:     %v
	DateTime: %v
	`, f.Name, f.CreatedAt.Format(time.RFC3339))
}
