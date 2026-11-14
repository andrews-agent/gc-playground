package inventory

import (
	"errors"
	"fmt"
	"slices"
	"testing"
)

func TestAddItemNormalizesSKU(t *testing.T) {
	for _, sku := range []string{"BOLT-10", "bolt-10", "Bolt-10"} {
		t.Run(sku, func(t *testing.T) {
			s := NewStore()
			s.AddWarehouse(Warehouse{ID: 1, Name: "Berlin"})
			item := Item{SKU: sku, Name: "Bolt 10 mm", WarehouseID: 1, Quantity: 20}
			if err := s.AddItem(item); err != nil {
				t.Fatal(err)
			}
			want := item
			want.SKU = "BOLT-10"
			for _, lookup := range []string{"BOLT-10", "bolt-10", "Bolt-10"} {
				got, err := s.Item(lookup)
				if err != nil {
					t.Errorf("Item(%q): %v", lookup, err)
				} else if got != want {
					t.Errorf("Item(%q) = %+v, want %+v", lookup, got, want)
				}
			}
			if got := s.Items(); len(got) != 1 || got[0] != want {
				t.Errorf("Items() = %+v, want [%+v]", got, want)
			}
		})
	}
}

func TestAddItemReplacesExistingSKU(t *testing.T) {
	for _, sku := range []string{"BOLT-10", "bolt-10", "Bolt-10"} {
		t.Run(sku, func(t *testing.T) {
			s := Sample()
			want := Item{SKU: "BOLT-10", Name: "Replacement bolt", WarehouseID: 2, Quantity: 7}
			before := s.Items()
			item := want
			item.SKU = sku
			if err := s.AddItem(item); err != nil {
				t.Fatal(err)
			}
			for i := range before {
				if before[i].SKU == want.SKU {
					before[i] = want
				}
			}
			if got := s.Items(); !slices.Equal(got, before) {
				t.Errorf("items after replacement = %+v, want %+v", got, before)
			}
		})
	}
}

func TestRejectedAddItemPreservesItems(t *testing.T) {
	for _, sku := range []string{"NEW", "BOLT-10", "bolt-10"} {
		t.Run(sku, func(t *testing.T) {
			s := Sample()
			before := s.Items()
			err := s.AddItem(Item{SKU: sku, Name: "Rejected item", WarehouseID: 99, Quantity: 7})
			if !errors.Is(err, ErrNotFound) {
				t.Errorf("AddItem error = %v, want ErrNotFound in chain", err)
			}
			if after := s.Items(); !slices.Equal(after, before) {
				t.Errorf("items after rejected add = %+v, want %+v", after, before)
			}
		})
	}
}

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

func TestReportsDiscardOutputOnError(t *testing.T) {
	for _, tt := range []struct {
		name string
		run  func(*Store) (string, error)
	}{
		{"show missing item", func(s *Store) (string, error) { return ItemLine(s, "MISSING") }},
		{"show missing warehouse", func(s *Store) (string, error) { return ItemLine(s, "GEAR-L") }},
		{"stock missing warehouse after valid rows", StockReport},
		{"low stock missing warehouse", func(s *Store) (string, error) { return LowStock(s, 5) }},
	} {
		t.Run(tt.name, func(t *testing.T) {
			s := Sample()
			// GEAR-L sorts after valid stock-report rows from other warehouses.
			delete(s.warehouses, 3)
			out, err := tt.run(s)
			if !errors.Is(err, ErrNotFound) {
				t.Errorf("error = %v, want ErrNotFound in chain", err)
			}
			if out != "" {
				t.Errorf("output on error = %q, want empty output", out)
			}
		})
	}
}

func TestReportsEmptyStore(t *testing.T) {
	for _, tt := range []struct {
		name string
		run  func(*Store) (string, error)
	}{
		{"stock", StockReport},
		{"low negative threshold", func(s *Store) (string, error) { return LowStock(s, -1) }},
		{"low zero threshold", func(s *Store) (string, error) { return LowStock(s, 0) }},
		{"low positive threshold", func(s *Store) (string, error) { return LowStock(s, 5) }},
	} {
		t.Run(tt.name, func(t *testing.T) {
			out, err := tt.run(NewStore())
			if err != nil {
				t.Fatal(err)
			}
			if out != "" {
				t.Errorf("empty-store report = %q, want empty output", out)
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
		t.Run(fmt.Sprint(tt.threshold), func(t *testing.T) {
			out, err := LowStock(Sample(), tt.threshold)
			if err != nil {
				t.Fatal(err)
			}
			if out != tt.want {
				t.Errorf("threshold %d: got %q, want %q", tt.threshold, out, tt.want)
			}
		})
	}
}

func TestLowStockSkipsUnselectedWarehouses(t *testing.T) {
	for _, tt := range []struct {
		name      string
		threshold int
		want      string
	}{
		{"no matches", -1, ""},
		{"some matches", 5, "GEAR-L\t0\tTbilisi\nGEAR-S\t5\tLisbon\nNUT-10\t4\tBerlin\n"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			s := Sample()
			s.AddWarehouse(Warehouse{ID: 4, Name: "Missing"})
			if err := s.AddItem(Item{SKU: "HIGH", WarehouseID: 4, Quantity: 100}); err != nil {
				t.Fatal(err)
			}
			delete(s.warehouses, 4)
			out, err := LowStock(s, tt.threshold)
			if err != nil {
				t.Fatal(err)
			}
			if out != tt.want {
				t.Errorf("LowStock = %q, want %q", out, tt.want)
			}
		})
	}
}

func TestLowStockQueryBudget(t *testing.T) {
	for _, threshold := range []int{-1, 0, 5, 200} {
		t.Run(fmt.Sprint(threshold), func(t *testing.T) {
			s := Sample()
			before := s.Queries
			if _, err := LowStock(s, threshold); err != nil {
				t.Fatal(err)
			}
			// One item query and at most one batch warehouse query,
			// regardless of how many items match or share a warehouse.
			if got := s.Queries - before; got > 2 {
				t.Errorf("LowStock query count = %d, want at most 2", got)
			}
		})
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

func TestRejectedMovePreservesItems(t *testing.T) {
	tests := []struct {
		name     string
		sku      string
		to       int
		quantity int
	}{
		{"missing item", "MISSING", 2, 1},
		{"missing warehouse", "BOLT-10", 99, 1},
		{"insufficient stock", "BOLT-10", 2, 121},
		{"negative quantity existing destination", "BOLT-10", 2, -1},
		{"negative quantity new destination", "BOLT-10", 3, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sample()
			if err := s.AddItem(Item{SKU: "BOLT-10@2", Name: "Bolt 10 mm", WarehouseID: 2, Quantity: 7}); err != nil {
				t.Fatal(err)
			}
			before := s.Items()
			if err := s.Move(tt.sku, tt.to, tt.quantity); err == nil {
				t.Error("Move succeeded, want an error")
			}
			if after := s.Items(); !slices.Equal(after, before) {
				t.Errorf("items after rejected move = %+v, want %+v", after, before)
			}
		})
	}
}

func TestMoveCaseInsensitiveSKU(t *testing.T) {
	for _, sku := range []string{"BOLT-10", "bolt-10", "Bolt-10"} {
		t.Run(sku, func(t *testing.T) {
			s := Sample()
			if err := s.Move(sku, 2, 20); err != nil {
				t.Fatal(err)
			}
			want := []Item{
				{SKU: "BELT-1", Name: "Drive belt", WarehouseID: 2, Quantity: 31},
				{SKU: "BOLT-10", Name: "Bolt 10 mm", WarehouseID: 1, Quantity: 100},
				{SKU: "BOLT-10@2", Name: "Bolt 10 mm", WarehouseID: 2, Quantity: 20},
				{SKU: "GEAR-L", Name: "Large gear", WarehouseID: 3, Quantity: 0},
				{SKU: "GEAR-S", Name: "Small gear", WarehouseID: 2, Quantity: 5},
				{SKU: "NUT-10", Name: "Nut 10 mm", WarehouseID: 1, Quantity: 4},
			}
			if got := s.Items(); !slices.Equal(got, want) {
				t.Errorf("items after Move(%q) = %+v, want %+v", sku, got, want)
			}
		})
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
	for _, quantity := range []int{20, 15, 0, 85, 0} {
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

func TestMoveDestinationCapacity(t *testing.T) {
	maxInt := int(^uint(0) >> 1)
	for _, tt := range []struct {
		name        string
		destination int
		quantity    int
		wantError   bool
	}{
		{"exact capacity", maxInt - 1, 1, false},
		{"past capacity", maxInt - 1, 2, true},
		{"full destination", maxInt, 1, true},
		{"zero to full destination", maxInt, 0, false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			s := Sample()
			destination := Item{SKU: "BOLT-10@2", Name: "Bolt 10 mm", WarehouseID: 2, Quantity: tt.destination}
			if err := s.AddItem(destination); err != nil {
				t.Fatal(err)
			}
			before := s.Items()
			err := s.Move("BOLT-10", 2, tt.quantity)
			if tt.wantError {
				if err == nil {
					t.Error("Move succeeded, want an overflow error")
				}
				if after := s.Items(); !slices.Equal(after, before) {
					t.Errorf("items after rejected move = %+v, want %+v", after, before)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			source, err := s.Item("BOLT-10")
			if err != nil {
				t.Fatal(err)
			}
			got, err := s.Item(destination.SKU)
			if err != nil {
				t.Fatal(err)
			}
			destination.Quantity += tt.quantity
			wantSource := Item{SKU: "BOLT-10", Name: "Bolt 10 mm", WarehouseID: 1, Quantity: 120 - tt.quantity}
			if source != wantSource {
				t.Errorf("source = %+v, want %+v", source, wantSource)
			}
			if got != destination {
				t.Errorf("destination = %+v, want %+v", got, destination)
			}
		})
	}
}

func TestMoveRejectsDestinationWarehouseConflict(t *testing.T) {
	for _, quantity := range []int{0, 20} {
		t.Run(fmt.Sprint(quantity), func(t *testing.T) {
			s := Sample()
			if err := s.AddItem(Item{SKU: "BOLT-10@2", Name: "Other stock", WarehouseID: 3, Quantity: 7}); err != nil {
				t.Fatal(err)
			}
			before := s.Items()
			if err := s.Move("BOLT-10", 2, quantity); err == nil {
				t.Error("Move succeeded with destination SKU in another warehouse")
			}
			if after := s.Items(); !slices.Equal(after, before) {
				t.Errorf("items after rejected move = %+v, want %+v", after, before)
			}
		})
	}
}
