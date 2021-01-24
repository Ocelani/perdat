package test

import (
	"os"
	"testing"

	"github.com/Ocelani/perdat/database"
	"gorm.io/gorm"
)

type TestDB struct {
	file     string
	result   *gorm.DB
	hasError bool
}

func Test_DBconnection(t *testing.T) {
	var (
		f = "perdat.db"
		i = TestDB{
			file:     f,
			hasError: false,
		}
	)
	result, err := database.Connect()
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
	if _, ok := r.(*gorm.DB); ok == true {
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
