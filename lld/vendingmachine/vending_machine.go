package main

import (
	"fmt"
	m "lld/vendingmachine/models"
	s "lld/vendingmachine/states"
)

type VendingMachine struct {
	Balance         int
	SelectedProduct string
	State           s.State
	Inventory       *m.Inventory
}

func NewVendingMachine() *VendingMachine {
	vm := &VendingMachine{
		Balance:   0,
		Inventory: m.NewInventory(),
	}

	vm.State = s.NewIdleState(vm)
	return vm
}

func (vm *VendingMachine) GetBalance() int {
	return vm.Balance
}

func (vm *VendingMachine) AddBalance(amount int) {
	if amount < 0 {
		fmt.Println("Amount cannot be negative")
		return
	}

	vm.Balance += amount
}

func (vm *VendingMachine) GetState() s.State {
	return vm.State
}

func (vm *VendingMachine) SetState(state s.State) {
	vm.State = state
}

func (vm *VendingMachine) SetSelectedProduct(code string) {
	vm.SelectedProduct = code
}

func (vm *VendingMachine) GetSelectedProduct() *m.Item {
	code := vm.SelectedProduct
	val, found := vm.Inventory.ItemShelves[code]
	if !found {
		fmt.Println("Something went wrong")
		return nil
	}

	return val.Item
}

func (vm *VendingMachine) GetProductByCode(code string) (*m.Item, bool) {
	v, ok := vm.Inventory.ItemShelves[code]
	if !ok {
		return nil, false
	}

	return v.Item, ok
}

func (vm *VendingMachine) AddToInventory(shelves map[string]*m.ItemShelf) {
	for key, shelf := range shelves {
		vm.Inventory.ItemShelves[key] = shelf
	}
}

func (vm *VendingMachine) DeleteFromInventory(code string) {
	_, exists := vm.Inventory.ItemShelves[code]
	if !exists {
		fmt.Println("Nothing to remove")
		return
	}

	delete(vm.Inventory.ItemShelves, code)
}

func (vm *VendingMachine) Dispense() *m.Item {
	code := vm.SelectedProduct

	product, found := vm.Inventory.ItemShelves[code]
	if !found {
		fmt.Println("Product not found")
		return nil
	}

	q := vm.Inventory.ItemShelves[code].Quantity
	q--

	if q <= 0 {
		vm.Inventory.ItemShelves[code].InStock = false
		vm.Inventory.ItemShelves[code].Quantity = 0
	}

	return product.Item
}
