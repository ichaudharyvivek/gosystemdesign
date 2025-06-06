package main

import (
	"fmt"
	m "lld/vendingmachine/models"
)

func main() {
	shelves := map[string]*m.ItemShelf{
		"101": m.NewItemShelf("101", 10, m.NewItem("Lay's Cream and Onion", m.CHIPS, 15)),
		"102": m.NewItemShelf("102", 5, m.NewItem("Pepsi Zero Sugar", m.SOFT_DRINKS, 5)),
		"103": m.NewItemShelf("103", 5, m.NewItem("Britannia Cream Biscuits", m.BISCUITS, 2)),
	}

	vm := NewVendingMachine()
	vm.AddInventory(shelves)

	fmt.Println("-- Adding Coins ---")
	vm.state.InsertCoins([]m.Coin{m.ONE})

	fmt.Println("-- Selecting Product --")
	vm.state.SelectProduct("103")

	fmt.Println("-- Dispensing Product --")
	vm.state.Dispense()
}
