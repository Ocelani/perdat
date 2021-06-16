package habit

import (
	"fmt"

	"github.com/Ocelani/perdat/pkg/entity"
)

type UpdateHabitNames map[string]string

// MakeMapUpdate returns a new UpdateHabitNames from a list of old, new string
// pairs. Replacements are performed in the order they appear in the
// target string, without overlapping matches. The old string
// comparisons are done in argument order.
func MakeMapUpdate(oldnew []string) UpdateHabitNames {
	update := UpdateHabitNames{}
	for i := 1; i <= len(oldnew); i++ {
		update[oldnew[i-1]] = oldnew[i]
	}
	return update
}

// ExcludeExistingHabits compares both slices according to its habit names
// and returns a new slice of nonexisting habits from the first slice parameter.
func ExcludeExistingHabits(habits, compare *[]entity.Habit) (*[]entity.Habit, int) {
	var (
		changed int
		exists  = []entity.Habit{}
	)
	for _, habit := range *habits {
		for _, comp := range *compare {
			if habit.Name == comp.Name {
				continue
			}
			exists = append(exists, habit)
			changed++
		}
	}
	return &exists, changed
}

func validateHabits(t *[]entity.Habit) error {
	if t == nil || len(*t) == 0 {
		return fmt.Errorf("invalid slice of habits")
	}
	return nil
}
