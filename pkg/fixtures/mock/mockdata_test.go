package mock

import "testing"

func TestMockActiveAlerts(t *testing.T) {
	r := MockActiveAlerts()
	if len(r) <= 0 {
		t.Fatalf("no data")
	}
}
