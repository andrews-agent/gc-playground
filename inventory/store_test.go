package inventory

import (
	"errors"
	"maps"
	"testing"
)

func TestWarehousesMissingID(t *testing.T) {
	for _, tt := range []struct {
		name string
		ids  []int
	}{
		{"only ID", []int{99}},
		{"first ID", []int{99, 1, 2}},
		{"middle ID", []int{1, 99, 2}},
		{"last ID", []int{1, 2, 99}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Sample().Warehouses(tt.ids)
			if !errors.Is(err, ErrNotFound) {
				t.Errorf("Warehouses(%v) error = %v, want ErrNotFound in chain", tt.ids, err)
			}
		})
	}
}

func TestWarehouse(t *testing.T) {
	for _, want := range []Warehouse{
		{ID: 1, Name: "Berlin"},
		{ID: 2, Name: "Lisbon"},
		{ID: 3, Name: "Tbilisi"},
	} {
		t.Run(want.Name, func(t *testing.T) {
			got, err := Sample().Warehouse(want.ID)
			if err != nil {
				t.Fatal(err)
			}
			if got != want {
				t.Errorf("Warehouse(%d) = %+v, want %+v", want.ID, got, want)
			}
		})
	}
}

func TestWarehouses(t *testing.T) {
	berlin := Warehouse{ID: 1, Name: "Berlin"}
	lisbon := Warehouse{ID: 2, Name: "Lisbon"}
	tests := []struct {
		name string
		ids  []int
		want map[int]Warehouse
	}{
		{"nil", nil, map[int]Warehouse{}},
		{"empty", []int{}, map[int]Warehouse{}},
		{"single", []int{2}, map[int]Warehouse{2: lisbon}},
		{"multiple", []int{2, 1}, map[int]Warehouse{1: berlin, 2: lisbon}},
		{"duplicates", []int{2, 1, 2}, map[int]Warehouse{1: berlin, 2: lisbon}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sample()
			got, err := s.Warehouses(tt.ids)
			if err != nil {
				t.Fatal(err)
			}
			if !maps.Equal(got, tt.want) {
				t.Errorf("Warehouses(%v) = %v, want %v", tt.ids, got, tt.want)
			}
		})
	}
}

func TestStoreQueries(t *testing.T) {
	tests := []struct {
		name string
		run  func(*Store)
		want int
	}{
		{"warehouse", func(s *Store) { s.Warehouse(1) }, 1},
		{"missing warehouse", func(s *Store) { s.Warehouse(99) }, 1},
		{"warehouses", func(s *Store) { s.Warehouses([]int{1, 2, 3}) }, 1},
		{"duplicate warehouses", func(s *Store) { s.Warehouses([]int{1, 1}) }, 1},
		{"empty warehouses", func(s *Store) { s.Warehouses(nil) }, 1},
		{"missing warehouses", func(s *Store) { s.Warehouses([]int{1, 99}) }, 1},
		{"item", func(s *Store) { s.Item("BOLT-10") }, 1},
		{"missing item", func(s *Store) { s.Item("MISSING") }, 1},
		{"items", func(s *Store) { s.Items() }, 1},
		{"add warehouse", func(s *Store) { s.AddWarehouse(Warehouse{ID: 4}) }, 0},
		{"add item", func(s *Store) { s.AddItem(Item{SKU: "NEW", WarehouseID: 1}) }, 1},
		{"add item missing warehouse", func(s *Store) { s.AddItem(Item{WarehouseID: 99}) }, 1},
		{"move", func(s *Store) { s.Move("BOLT-10", 2, 1) }, 2},
		{"move missing item", func(s *Store) { s.Move("MISSING", 2, 1) }, 1},
		{"move missing warehouse", func(s *Store) { s.Move("BOLT-10", 99, 1) }, 2},
		{"move insufficient stock", func(s *Store) { s.Move("BOLT-10", 2, 121) }, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sample()
			before := s.Queries
			tt.run(s)
			if got := s.Queries - before; got != tt.want {
				t.Errorf("query count = %d, want %d", got, tt.want)
			}
		})
	}
}
