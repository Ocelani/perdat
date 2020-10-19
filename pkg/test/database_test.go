package test

import (
	"testing"

	"github.com/Ocelani/perdat/database"
)

type TestItem struct {
	file     string
	result   *database.DB
	hasError bool
}

func TestConnection(t *testing.T) {
	var i = TestItem{
		file: "perdat.db",
	}

	result, err := database.Connect(i.file)
	if err != nil {
		t.Errorf("\nFAILED!\n☐ Connect(%v) \n- Expected: %v,\n✘ Got: %v \n", i.file, true, err)
	}

	var r interface{} = result
	_, ok := r.(*database.DB)

	if ok == true {
		t.Logf("\nPASSED!\n☐ Connect(%v) \n- Expected: %v \n✔ Got: %v \n", i.file, true, ok)
	} else {
		t.Errorf("\nFAILED!\n☐ Connect(%v) \n- Expected: %v,\n✘ Got: %v \n", i.file, true, ok)
	}
}
