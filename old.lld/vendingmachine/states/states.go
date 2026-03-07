package states

import m "lld/vendingmachine/models"

type Machine interface {
	AddBalance(amount int)
	GetBalance() int

	SetState(state State)
	GetState() State

	GetProductByCode(code string) (*m.Item, bool)
	IsProductOutOfStock(code string) bool

	SetSelectedProduct(code string)
	GetSelectedProduct() *m.Item

	DispenseProduct() (int, bool)
	Refund(amount int)
}

type State interface {
	InsertCoins(coins []m.Coin)
	SelectProduct(code string)
	Dispense()
}
