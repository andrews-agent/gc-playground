package inventory

import (
	"errors"
	"testing"
)

func TestNotFoundWrapping(t *testing.T) {
	tests := []struct {
		name string
		run  func(*Store) error
	}{
		{"add item warehouse", func(s *Store) error {
			return s.AddItem(Item{SKU: "MISSING", WarehouseID: 99})
		}},
		{"move item", func(s *Store) error { return s.Move("MISSING", 1, 1) }},
		{"move warehouse", func(s *Store) error { return s.Move("BOLT-10", 99, 1) }},
		{"show item", func(s *Store) error { _, err := ItemLine(s, "MISSING"); return err }},
		{"show warehouse", func(s *Store) error {
			delete(s.warehouses, 1)
			_, err := ItemLine(s, "BOLT-10")
			return err
		}},
		{"stock warehouse", func(s *Store) error {
			delete(s.warehouses, 1)
			_, err := StockReport(s)
			return err
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.run(Sample()); !errors.Is(err, ErrNotFound) {
				t.Errorf("error = %v, want ErrNotFound in chain", err)
			}
		})
	}
}

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
