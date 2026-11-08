package inventory

import (
	"strings"
	"testing"
)

func TestStockReport(t *testing.T) {
	out, err := StockReport(Sample())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)
}

func TestLowStock(t *testing.T) {
	tests := []struct {
		threshold int
		want      []string
	}{
		{threshold: 1, want: []string{"GEAR-L"}},
		{threshold: 5, want: []string{"GEAR-L", "GEAR-S", "NUT-10"}},
		{threshold: 200, want: []string{"BELT-1", "BOLT-10", "GEAR-L", "GEAR-S", "NUT-10"}},
	}
	for _, tt := range tests {
		out, err := LowStock(Sample(), tests[0].threshold)
		if err != nil {
			t.Fatal(err)
		}
		for _, sku := range tests[0].want {
			if !strings.Contains(out, sku) {
				t.Errorf("threshold %d: %s missing from\n%s", tt.threshold, sku, out)
			}
		}
	}
}

func TestItemLine(t *testing.T) {
	line, err := ItemLine(Sample(), "gear-s")
	if err != nil {
		t.Fatal(err)
	}
	if line != "GEAR-S Small gear: 5 in Lisbon" {
		t.Errorf("got %q", line)
	}
}

func TestMoveInBackground(t *testing.T) {
	s := Sample()
	done := make(chan error)
	go func() {
		done <- s.Move("BOLT-10", 2, 20)
	}()
	if err := <-done; err != nil {
		t.Fatal(err)
	}
	item, err := s.Item("BOLT-10")
	if err != nil {
		t.Fatal(err)
	}
	if item.Quantity != 100 {
		t.Errorf("got %d", item.Quantity)
	}
}
