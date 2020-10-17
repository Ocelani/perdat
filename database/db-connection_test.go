package database

import (
	"testing"

	"xorm.io/xorm"
)

type Main interface {
	main()
}

type Connection struct {
}

// TestItem type
type TestItem struct {
	inputs   Connection
	result   *xorm.Engine
	hasError bool
}

func testExecute(item TestItem, t *testing.T) {
	// get result of func()
	_, err := run()

	// expected an error
	if err == nil {
		t.Errorf("\nFAILED!\n☐ Intersection()\n   \n- Expected: an error,\n✘ Got: value ")
	} else {
		t.Logf("\nPASSED!\n☐ Intersection()\n   \n- Expected: an error,\n✔ Got: an error \n")
	}
}
