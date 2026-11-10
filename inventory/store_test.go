package inventory

import "testing"

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
