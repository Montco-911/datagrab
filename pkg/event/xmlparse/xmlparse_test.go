package xmlparse

import (
	"github.com/Montco-911/datagrab/pkg/fixtures/mock"
	"testing"
)

func TestDecodeXML(t *testing.T) {


	d := []byte(mock.MockActiveAlerts())
	a := Decode(d)
	if len(a.Events) != 13 {
		t.Fatalf(("Didn't get all records"))
	}
	expected := "EMS: GENERAL WEAKNESS"
	if a.Events[3].Title != expected {
		t.Fatalf("Expected: %s, got %s\n", expected, a.Events[3].Title)
	}

}
