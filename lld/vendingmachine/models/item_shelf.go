package models

type Item struct {
	ItemType ItemType
	Name     string
	Price    int
}

func NewItem(iType ItemType, name string, price int) *Item {
	return &Item{
		ItemType: iType,
		Name:     name,
		Price:    price,
	}
}

type ItemShelf struct {
	Quantity int
	InStock  bool
	Code     string
	Item     *Item
}

func NewItemShelf(q int, code string, item *Item) *ItemShelf {
	return &ItemShelf{
		Quantity: q,
		InStock:  true,
		Code:     code,
		Item:     item,
	}
}
