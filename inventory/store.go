// Package inventory keeps items in warehouses and reports on their stock.
package inventory

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// ErrNotFound is returned when an item or a warehouse does not exist.
var ErrNotFound = errors.New("not found")

type Warehouse struct {
	ID   int
	Name string
}

type Item struct {
	SKU         string
	Name        string
	WarehouseID int
	Quantity    int
}

// Store is an in-memory stand-in for the database.
// Queries counts calls to Warehouse, Warehouses, Item, and Items, including
// failed lookups and calls made internally by other methods.
type Store struct {
	warehouses map[int]Warehouse
	items      map[string]Item
	Queries    int
}

func NewStore() *Store {
	return &Store{warehouses: map[int]Warehouse{}, items: map[string]Item{}}
}

func (s *Store) AddWarehouse(w Warehouse) {
	s.warehouses[w.ID] = w
}

// AddItem stores an item with its SKU normalized to uppercase for lookup.
func (s *Store) AddItem(item Item) error {
	if _, err := s.Warehouse(item.WarehouseID); err != nil {
		return fmt.Errorf("add item %s: %w", item.SKU, err)
	}
	item.SKU = strings.ToUpper(item.SKU)
	s.items[item.SKU] = item
	return nil
}

// Warehouse returns one warehouse by its ID.
func (s *Store) Warehouse(id int) (Warehouse, error) {
	s.Queries++
	w, ok := s.warehouses[id]
	if !ok {
		return Warehouse{}, fmt.Errorf("warehouse %d: %w", id, ErrNotFound)
	}
	return w, nil
}

// Warehouses returns the warehouses with the given IDs in one query.
func (s *Store) Warehouses(ids []int) (map[int]Warehouse, error) {
	s.Queries++
	found := make(map[int]Warehouse, len(ids))
	for _, id := range ids {
		w, ok := s.warehouses[id]
		if !ok {
			return nil, fmt.Errorf("warehouse %d: %w", id, ErrNotFound)
		}
		found[id] = w
	}
	return found, nil
}

// Item returns one item by its SKU.
func (s *Store) Item(sku string) (Item, error) {
	s.Queries++
	item, ok := s.items[strings.ToUpper(sku)]
	if !ok {
		return Item{}, fmt.Errorf("item %s: %w", sku, ErrNotFound)
	}
	return item, nil
}

// Items returns every item, sorted by SKU.
func (s *Store) Items() []Item {
	s.Queries++
	items := make([]Item, 0, len(s.items))
	for _, item := range s.items {
		items = append(items, item)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].SKU < items[j].SKU })
	return items
}

// Move moves quantity of an item to another warehouse.
func (s *Store) Move(sku string, to int, quantity int) error {
	item, err := s.Item(sku)
	if err != nil {
		return fmt.Errorf("move: %w", err)
	}
	if _, err := s.Warehouse(to); err != nil {
		return fmt.Errorf("move: %w", err)
	}
	if quantity > item.Quantity {
		return fmt.Errorf("move %s: only %d left", sku, item.Quantity)
	}
	item.Quantity -= quantity
	s.items[item.SKU] = item
	moved := item
	moved.SKU = item.SKU + "@" + fmt.Sprint(to)
	moved.WarehouseID = to
	moved.Quantity = s.items[moved.SKU].Quantity + quantity
	s.items[moved.SKU] = moved
	return nil
}
