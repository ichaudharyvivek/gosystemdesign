package models

type Inventory struct {
	Shelves map[string]*ItemShelf
}

func NewInventory(shelves map[string]*ItemShelf) *Inventory {
	return &Inventory{shelves}
}
