package test

import (
	"os"
	"testing"

	"github.com/Ocelani/perdat/database"
)

type TestItem struct {
	file     string
	result   *database.DB
	hasError bool
}

func Test_DBconnection(t *testing.T) {
	var (
		f = "perdat.db"
		i = TestItem{
			file:     f,
			hasError: false,
		}
	)
	result, err := database.Connect(
		i.file, database.NewStdoutLogger())
	if err != nil {
		t.Errorf(`
			FAILED!
			☐ Connect(%v)
			- Expected: %v
			✘ Got: %v
			`, i.file, true, err,
		)
	}
	var r interface{} = result
	if _, ok := r.(*database.DB); ok == true {
		t.Logf(`
			PASSED!
			☐ Connect(%v)
			- Expected: %v
			✔ Got: %v
			`, i.file, true, ok,
		)
	} else {
		t.Errorf(`
			FAILED!
			☐ Connect(%v) 
			- Expected: %v,
			✘ Got: %v 
			`, i.file, true, ok,
		)
	}
	os.Remove(f)
}
