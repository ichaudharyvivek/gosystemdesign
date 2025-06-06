package models

type Item struct {
	ItemType ItemType
	Name     string
	Price    int
}

type ItemShelf struct {
	Quantity     int
	IsOutOfStock bool
	Code         string
	Item         *Item
}
