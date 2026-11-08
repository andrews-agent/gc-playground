package inventory

import (
	"fmt"
	"strings"
)

// StockReport lists every item with the name of its warehouse.
func StockReport(s *Store) (string, error) {
	items := s.Items()
	if len(items) == 0 {
		return "", nil
	}
	ids := make([]int, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.WarehouseID)
	}
	warehouses, err := s.Warehouses(ids)
	if err != nil {
		return "", fmt.Errorf("stock report: %v", err)
	}
	var out strings.Builder
	for _, item := range items {
		w := warehouses[item.WarehouseID]
		fmt.Fprintf(&out, "%s\t%s\t%d\t%s\n", item.SKU, item.Name, item.Quantity, w.Name)
	}
	return out.String(), nil
}

// LowStock lists the items whose quantity is at or below the threshold,
// with the name of their warehouse.
func LowStock(s *Store, threshold int) (string, error) {
	var low []Item
	var ids []int
	for _, item := range s.Items() {
		if item.Quantity <= threshold {
			low = append(low, item)
			ids = append(ids, item.WarehouseID)
		}
	}
	warehouses, err := s.Warehouses(ids)
	if err != nil {
		return "", fmt.Errorf("low stock: %w", err)
	}
	var out strings.Builder
	for _, item := range low {
		fmt.Fprintf(&out, "%s\t%d\t%s\n", item.SKU, item.Quantity, warehouses[item.WarehouseID].Name)
	}
	return out.String(), nil
}

// ItemLine describes one item, for the show command.
func ItemLine(s *Store, sku string) (string, error) {
	item, err := s.Item(sku)
	if err != nil {
		return "", fmt.Errorf("show: %v", err)
	}
	w, err := s.Warehouse(item.WarehouseID)
	if err != nil {
		return "", fmt.Errorf("show: %v", err)
	}
	return fmt.Sprintf("%s %s: %d in %s", item.SKU, item.Name, item.Quantity, w.Name), nil
}
