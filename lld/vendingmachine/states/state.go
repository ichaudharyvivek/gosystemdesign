package states

import m "lld/vendingmachine/models"

// Vending Machine
type Machine interface {
	GetBalance() int
	AddBalance(amount int)

	GetState() State
	SetState(state State)

	GetSelectedProduct() *m.Item
	SetSelectedProduct(code string)

	AddToInventory(shelves map[string]*m.ItemShelf)
	RemoveFromInventory(code string)

	Dispense() *m.Item
}

// Behaviour of VM
type State interface {
	InsertCoins(coins []m.Coin)
	SelectProduct(code string)
	Dispense()
	OutOfOrder()
}
