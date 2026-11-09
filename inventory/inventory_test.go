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
		{"low stock warehouse", func(s *Store) error {
			delete(s.warehouses, 1)
			_, err := LowStock(s, 5)
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
	want := "BELT-1\tDrive belt\t31\tLisbon\n" +
		"BOLT-10\tBolt 10 mm\t120\tBerlin\n" +
		"GEAR-L\tLarge gear\t0\tTbilisi\n" +
		"GEAR-S\tSmall gear\t5\tLisbon\n" +
		"NUT-10\tNut 10 mm\t4\tBerlin\n"
	if out != want {
		t.Errorf("StockReport = %q, want %q", out, want)
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
	wantSource := Item{SKU: "BOLT-10", Name: "Bolt 10 mm", WarehouseID: 1, Quantity: 100}
	if item != wantSource {
		t.Errorf("source item = %+v, want %+v", item, wantSource)
	}
	moved, err := s.Item("BOLT-10@2")
	if err != nil {
		t.Fatal(err)
	}
	wantDestination := Item{SKU: "BOLT-10@2", Name: "Bolt 10 mm", WarehouseID: 2, Quantity: 20}
	if moved != wantDestination {
		t.Errorf("destination item = %+v, want %+v", moved, wantDestination)
	}
}

func TestMoveAccumulatesDestination(t *testing.T) {
	s := Sample()
	transferred := 0
	for _, quantity := range []int{20, 15, 0} {
		if err := s.Move("BOLT-10", 2, quantity); err != nil {
			t.Fatal(err)
		}
		transferred += quantity
		source, err := s.Item("BOLT-10")
		if err != nil {
			t.Fatal(err)
		}
		destination, err := s.Item("BOLT-10@2")
		if err != nil {
			t.Fatal(err)
		}
		wantSource := Item{SKU: "BOLT-10", Name: "Bolt 10 mm", WarehouseID: 1, Quantity: 120 - transferred}
		wantDestination := Item{SKU: "BOLT-10@2", Name: "Bolt 10 mm", WarehouseID: 2, Quantity: transferred}
		if source != wantSource {
			t.Errorf("after moving %d: source = %+v, want %+v", quantity, source, wantSource)
		}
		if destination != wantDestination {
			t.Errorf("after moving %d: destination = %+v, want %+v", quantity, destination, wantDestination)
		}
	}
}
