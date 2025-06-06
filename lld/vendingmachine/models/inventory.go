package models

type Inventory struct {
	ItemShelves map[string]*ItemShelf
}

func NewInventory() *Inventory {
	return &Inventory{
		ItemShelves: make(map[string]*ItemShelf),
	}
}
