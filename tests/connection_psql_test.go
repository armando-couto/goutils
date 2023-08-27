package tests

import (
	"github.com/armando-couto/goutils"
	"testing"
)

func TestConnectionBDPostgreSQL(t *testing.T) {
	db := goutils.ConnectionBDPostgreSQL("Teste")
	t.Log(db.Stats())
}
