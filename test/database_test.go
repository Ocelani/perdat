package test

import (
	"os"
	"testing"

	"github.com/Ocelani/perdat/database"
)

func Test_DBconnection(t *testing.T) {
	file := "perdat.db"
	got, err := database.Connect()

	os.Remove(f)
}
