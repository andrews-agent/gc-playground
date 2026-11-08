package inventory

import (
	"testing"
)

func TestStockReport(t *testing.T) {
	s := Sample()
	s.Queries = 0
	out, err := StockReport(s)
	if err != nil {
		t.Fatal(err)
	}
	want := "BELT-1\tDrive belt\t31\tLisbon\nBOLT-10\tBolt 10 mm\t120\tBerlin\nGEAR-L\tLarge gear\t0\tTbilisi\nGEAR-S\tSmall gear\t5\tLisbon\nNUT-10\tNut 10 mm\t4\tBerlin\n"
	if out != want {
		t.Errorf("report = %q, want %q", out, want)
	}
	if s.Queries != 2 {
		t.Errorf("queries = %d, want 2", s.Queries)
	}
}

func TestStockReportEmpty(t *testing.T) {
	s := NewStore()
	out, err := StockReport(s)
	if err != nil || out != "" {
		t.Fatalf("StockReport(empty) = %q, %v", out, err)
	}
	if s.Queries != 1 {
		t.Errorf("queries = %d, want 1", s.Queries)
	}
}

func TestStockReportMissingWarehouse(t *testing.T) {
	s := Sample()
	delete(s.warehouses, 1)
	out, err := StockReport(s)
	if err == nil || err.Error() != "stock report: warehouse 1: not found" {
		t.Fatalf("error = %v, want missing warehouse error", err)
	}
	if out != "" {
		t.Errorf("report = %q, want no partial output", out)
	}
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
