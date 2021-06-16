package fact

import (
	"fmt"

	"github.com/Ocelani/perdat/pkg/entity"
)

type UpdateFactNames map[string]string

// MakeMapUpdate returns a new UpdateFactNames from a list of old, new string
// pairs. Replacements are performed in the order they appear in the
// target string, without overlapping matches. The old string
// comparisons are done in argument order.
func MakeMapUpdate(oldnew []string) UpdateFactNames {
	update := UpdateFactNames{}
	for i := 1; i <= len(oldnew); i++ {
		update[oldnew[i-1]] = oldnew[i]
	}
	return update
}

// ExcludeExistingFacts compares both slices according to its fact names
// and returns a new slice of nonexisting facts from the first slice parameter.
func ExcludeExistingFacts(facts, compare *[]entity.Fact) (*[]entity.Fact, int) {
	var (
		changed int
		exists  = []entity.Fact{}
	)
	for _, fact := range *facts {
		for _, comp := range *compare {
			if fact.Name == comp.Name {
				continue
			}
			exists = append(exists, fact)
			changed++
		}
	}
	return &exists, changed
}

func validateFacts(t *[]entity.Fact) error {
	if t == nil || len(*t) == 0 {
		return fmt.Errorf("invalid slice of facts")
	}
	return nil
}
