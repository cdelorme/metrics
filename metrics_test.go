package metrics

import "testing"

func TestNewStats(t *testing.T) {
	s := NewStats()
	if s == nil {
		t.Error("failed to create a new stats...")
	}
}
