package models

import "fmt"

type Item struct {
	Name     string
	ItemType ItemType
	Price    int
}

func (r Item) String() string {
	return fmt.Sprintf("Name:%s, Item Type:%s, Price:%d", r.Name, r.ItemType, r.Price)
}

func NewItem(name string, itemType ItemType, price int) *Item {
	return &Item{name, itemType, price}
}

type ItemShelf struct {
	Item     *Item
	Quantity int
	Code     string
	InStock  bool
}

func NewItemShelf(code string, quantity int, item *Item) *ItemShelf {
	return &ItemShelf{item, quantity, code, true}
}

func (r ItemShelf) String() string {
	return fmt.Sprintf("Code:%s, Quatity:%d, InStock:%t, \nItem: %s", r.Code, r.Quantity, r.InStock, r.Item)
}
