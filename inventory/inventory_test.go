package inventory

import (
	"testing"
	"time"
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
		want      string
	}{
		{threshold: -1, want: ""},
		{threshold: 0, want: "GEAR-L\t0\tTbilisi\n"},
		{threshold: 1, want: "GEAR-L\t0\tTbilisi\n"},
		{threshold: 5, want: "GEAR-L\t0\tTbilisi\nGEAR-S\t5\tLisbon\nNUT-10\t4\tBerlin\n"},
		{threshold: 200, want: "BELT-1\t31\tLisbon\nBOLT-10\t120\tBerlin\nGEAR-L\t0\tTbilisi\nGEAR-S\t5\tLisbon\nNUT-10\t4\tBerlin\n"},
	}
	for _, tt := range tests {
		out, err := LowStock(Sample(), tt.threshold)
		if err != nil {
			t.Fatal(err)
		}
		if out != tt.want {
			t.Errorf("threshold %d: got %q, want %q", tt.threshold, out, tt.want)
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
	go s.Move("BOLT-10", 2, 20)
	time.Sleep(10 * time.Millisecond)
	item, _ := s.Item("BOLT-10")
	if item.Quantity != 100 {
		t.Errorf("got %d", item.Quantity)
	}
}
