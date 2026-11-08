package inventory

// Sample returns a store with a few warehouses and items, for the command and the tests.
func Sample() *Store {
	s := NewStore()
	s.AddWarehouse(Warehouse{ID: 1, Name: "Berlin"})
	s.AddWarehouse(Warehouse{ID: 2, Name: "Lisbon"})
	s.AddWarehouse(Warehouse{ID: 3, Name: "Tbilisi"})
	for _, item := range []Item{
		{SKU: "BOLT-10", Name: "Bolt 10 mm", WarehouseID: 1, Quantity: 120},
		{SKU: "NUT-10", Name: "Nut 10 mm", WarehouseID: 1, Quantity: 4},
		{SKU: "GEAR-S", Name: "Small gear", WarehouseID: 2, Quantity: 5},
		{SKU: "GEAR-L", Name: "Large gear", WarehouseID: 3, Quantity: 0},
		{SKU: "BELT-1", Name: "Drive belt", WarehouseID: 2, Quantity: 31},
	} {
		if err := s.AddItem(item); err != nil {
			panic(err)
		}
	}
	return s
}
