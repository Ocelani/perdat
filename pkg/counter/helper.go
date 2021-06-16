package counter

import (
	"fmt"

	"github.com/Ocelani/perdat/pkg/entity"
)

type UpdateCounterNames map[string]string

// MakeMapUpdate returns a new UpdateCounterNames from a list of old, new string
// pairs. Replacements are performed in the order they appear in the
// target string, without overlapping matches. The old string
// comparisons are done in argument order.
func MakeMapUpdate(oldnew []string) UpdateCounterNames {
	update := UpdateCounterNames{}
	for i := 1; i <= len(oldnew); i++ {
		update[oldnew[i-1]] = oldnew[i]
	}
	return update
}

// ExcludeExistingCounters compares both slices according to its counter names
// and returns a new slice of nonexisting counters from the first slice parameter.
func ExcludeExistingCounters(counters, compare *[]entity.Counter) (*[]entity.Counter, int) {
	var (
		changed int
		exists  = []entity.Counter{}
	)
	for _, counter := range *counters {
		for _, comp := range *compare {
			if counter.Name == comp.Name {
				continue
			}
			exists = append(exists, counter)
			changed++
		}
	}
	return &exists, changed
}

func validateCounters(t *[]entity.Counter) error {
	if t == nil || len(*t) == 0 {
		return fmt.Errorf("invalid slice of counters")
	}
	return nil
}
