package database

import (
	"testing"

	"xorm.io/xorm"
)

type Connection interface {
	Connection()
}

type TestItem struct {
	file     string
	result   *xorm.Engine
	hasError bool
}

func TestConnection(t *testing.T) {
	var i = TestItem{
		file: "perdat.db",
	}
	result, err := Connect(i.file)

	var r interface{} = result
	_, ok := r.(*xorm.Engine)

	if err != nil {
		t.Errorf("\nFAILED!\n☐ Connect(%v) \n- Expected: %v,\n✘ Got: %v \n", i.file, true, err)
	}
	if ok == true {
		t.Logf("\nPASSED!\n☐ Connect(%v) \n- Expected: %v \n✔ Got: %v \n", i.file, true, ok)
	} else {
		t.Errorf("\nFAILED!\n☐ Connect(%v) \n- Expected: %v,\n✘ Got: %v \n", i.file, true, ok)
	}
}
