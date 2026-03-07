package states

import m "lld/coffeemachine/models"

type Machine interface {
	AddBalance(amount int)
	GetBalance() int

	SelectCoffee(id string) error
	GetSelectedCoffee() *m.Coffee

	SetState(st State)
	GetState() State

	BrewCoffee() (int, error)
}

type State interface {
	InsertMoney(amount int)
	Select(id string)
	Brew()
}
