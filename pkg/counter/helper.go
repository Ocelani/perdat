package counter

import (
	"fmt"
	"strconv"

	"github.com/Ocelani/perdat/pkg/entity"
)

type (
	UpdateCounterNames map[string]string
	UpdateCounterInt   map[string]int
)

// MakeMapUpdate returns a new UpdateCounterNames from a list of old, new string
// pairs. Replacements are performed in the order they appear in the
// target string, without overlapping matches. The old string
// comparisons are done in argument order.
func MakeMapUpdateNames(oldnew []string) UpdateCounterNames {
	update := UpdateCounterNames{}
	for i := 1; i <= len(oldnew); i++ {
		update[oldnew[i-1]] = oldnew[i]
	}
	return update
}

func MakeMapUpdateInt(args []string) UpdateCounterInt {
	setNums := UpdateCounterInt{}

	for i, a := range args {
		n, err := strconv.Atoi(a)

		if err != nil {
			setNums[a] = 1
		} else {
			name := args[i-1]
			setNums[name] = n
		}
	}

	return setNums
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
