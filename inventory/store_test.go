package inventory

import (
	"errors"
	"fmt"
	"maps"
	"slices"
	"testing"
)

func TestMissingLookupsInEmptyStore(t *testing.T) {
	for _, tt := range []struct {
		name string
		run  func(*testing.T, *Store) error
	}{
		{"item", func(t *testing.T, s *Store) error {
			_, err := s.Item("MISSING")
			return err
		}},
		{"warehouse", func(t *testing.T, s *Store) error {
			_, err := s.Warehouse(99)
			return err
		}},
		{"warehouses", func(t *testing.T, s *Store) error {
			got, err := s.Warehouses([]int{99})
			if got != nil {
				t.Errorf("missing warehouses = %v, want nil", got)
			}
			return err
		}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.run(t, NewStore()); !errors.Is(err, ErrNotFound) {
				t.Errorf("empty-store lookup error = %v, want ErrNotFound in chain", err)
			}
		})
	}
}

func TestMissingSingleLookupResults(t *testing.T) {
	for _, store := range []struct {
		name string
		new  func() *Store
	}{
		{"empty", NewStore},
		{"populated", Sample},
	} {
		t.Run(store.name, func(t *testing.T) {
			s := store.new()
			t.Run("item", func(t *testing.T) {
				got, err := s.Item("MISSING")
				if !errors.Is(err, ErrNotFound) {
					t.Errorf("Item error = %v, want ErrNotFound in chain", err)
				}
				if got != (Item{}) {
					t.Errorf("missing item = %+v, want zero value", got)
				}
			})
			t.Run("warehouse", func(t *testing.T) {
				got, err := s.Warehouse(99)
				if !errors.Is(err, ErrNotFound) {
					t.Errorf("Warehouse error = %v, want ErrNotFound in chain", err)
				}
				if got != (Warehouse{}) {
					t.Errorf("missing warehouse = %+v, want zero value", got)
				}
			})
		})
	}
}

func TestItems(t *testing.T) {
	a := Item{SKU: "A", Name: "First", WarehouseID: 1, Quantity: 7}
	b := Item{SKU: "B", Name: "Second", WarehouseID: 2, Quantity: 0}
	c := Item{SKU: "C", Name: "Third", WarehouseID: 1, Quantity: 12}
	for _, tt := range []struct {
		name  string
		items []Item
		want  []Item
	}{
		{"empty", nil, nil},
		{"single", []Item{b}, []Item{b}},
		{"out of order", []Item{c, a, b}, []Item{a, b, c}},
		{"reverse order", []Item{c, b, a}, []Item{a, b, c}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStore()
			s.AddWarehouse(Warehouse{ID: 1, Name: "Berlin"})
			s.AddWarehouse(Warehouse{ID: 2, Name: "Lisbon"})
			for _, item := range tt.items {
				if err := s.AddItem(item); err != nil {
					t.Fatal(err)
				}
			}
			if got := s.Items(); !slices.Equal(got, tt.want) {
				t.Errorf("Items() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestItemsResultIsIndependent(t *testing.T) {
	s := Sample()
	want := slices.Clone(s.Items())
	got := s.Items()
	got[0] = Item{SKU: "CHANGED", Quantity: -1}
	if stored := s.Items(); !slices.Equal(stored, want) {
		t.Errorf("items after changing result = %+v, want %+v", stored, want)
	}

	snapshot := s.Items()
	before := slices.Clone(snapshot)
	replacement := Item{SKU: "BELT-1", Name: "Replacement belt", WarehouseID: 2, Quantity: 7}
	if err := s.AddItem(replacement); err != nil {
		t.Fatal(err)
	}
	// Read again to catch implementations that reuse a result buffer.
	s.Items()
	if !slices.Equal(snapshot, before) {
		t.Errorf("previous result after replacing item = %+v, want %+v", snapshot, before)
	}
}

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
			got, err := Sample().Warehouses(tt.ids)
			if !errors.Is(err, ErrNotFound) {
				t.Errorf("Warehouses(%v) error = %v, want ErrNotFound in chain", tt.ids, err)
			}
			if got != nil {
				t.Errorf("Warehouses(%v) = %v, want nil on error", tt.ids, got)
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

func TestAddWarehouseReplacesExistingID(t *testing.T) {
	for _, id := range []int{1, 2, 3} {
		t.Run(fmt.Sprint(id), func(t *testing.T) {
			s := Sample()
			before, err := s.Warehouses([]int{1, 2, 3})
			if err != nil {
				t.Fatal(err)
			}
			want := Warehouse{ID: id, Name: "Replacement warehouse"}
			s.AddWarehouse(want)
			before[id] = want
			got, err := s.Warehouses([]int{1, 2, 3})
			if err != nil {
				t.Fatal(err)
			}
			if !maps.Equal(got, before) {
				t.Errorf("warehouses after replacing ID %d = %v, want %v", id, got, before)
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
			if got == nil {
				t.Fatal("Warehouses returned nil on success, want an initialized map")
			}
			if !maps.Equal(got, tt.want) {
				t.Errorf("Warehouses(%v) = %v, want %v", tt.ids, got, tt.want)
			}
		})
	}
}

func TestWarehousesResultIsIndependent(t *testing.T) {
	s := NewStore()
	want := Warehouse{ID: 1, Name: "Berlin"}
	s.AddWarehouse(want)
	got, err := s.Warehouses([]int{1})
	if err != nil {
		t.Fatal(err)
	}
	got[1] = Warehouse{ID: 1, Name: "Changed result"}
	stored, err := s.Warehouse(1)
	if err != nil {
		t.Fatal(err)
	}
	if stored != want {
		t.Errorf("warehouse after changing result = %+v, want %+v", stored, want)
	}
	delete(got, 1)
	stored, err = s.Warehouse(1)
	if err != nil {
		t.Fatal(err)
	}
	if stored != want {
		t.Errorf("warehouse after deleting result entry = %+v, want %+v", stored, want)
	}
	s.AddWarehouse(Warehouse{ID: 2, Name: "Lisbon"})
	if len(got) != 0 {
		t.Errorf("result after adding warehouse = %v, want empty map", got)
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
