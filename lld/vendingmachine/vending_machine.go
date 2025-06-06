package main

import (
	"fmt"
	m "lld/vendingmachine/models"
	s "lld/vendingmachine/states"
)

type VendingMachine struct {
	balance         int
	selectedProduct string
	inventory       *m.Inventory
	state           s.State
}

func NewVendingMachine() *VendingMachine {
	vm := &VendingMachine{}
	vm.state = s.NewIdleState(vm)
	return vm
}

func (vm VendingMachine) String() string {
	return fmt.Sprintf("Balance: %d, \nSelectedProduct: %s, \nInventory: %v, \nState: %v", vm.balance, vm.selectedProduct, vm.inventory, vm.state)
}

func (vm *VendingMachine) AddInventory(shelves map[string]*m.ItemShelf) {
	vm.inventory = m.NewInventory(shelves)
}

func (vm *VendingMachine) AddBalance(amount int) {
	if amount < 0 {
		fmt.Println("Amount cannot be negative")
		return
	}

	vm.balance += amount
}

func (vm *VendingMachine) GetBalance() int {
	return vm.balance
}

func (vm *VendingMachine) SetState(state s.State) {
	vm.state = state
}

func (vm *VendingMachine) GetState() s.State {
	return vm.state
}

func (vm *VendingMachine) GetProductByCode(code string) (*m.Item, bool) {
	prod, found := vm.inventory.Shelves[code]
	if !found {
		return nil, false
	}

	return prod.Item, true
}

func (vm *VendingMachine) SetSelectedProduct(code string) {
	vm.selectedProduct = code
}

func (vm *VendingMachine) GetSelectedProduct() *m.Item {
	prod, _ := vm.GetProductByCode(vm.selectedProduct)
	return prod
}

func (vm *VendingMachine) IsProductOutOfStock(code string) bool {
	return vm.inventory.Shelves[code].Quantity > 0
}

func (vm *VendingMachine) DispenseProduct() (int, bool) {
	code := vm.selectedProduct
	prod, _ := vm.GetProductByCode(code)

	vm.inventory.Shelves[code].Quantity -= 1
	refund := vm.balance - prod.Price
	if refund > 0 {
		return refund, false
	} else {
		return refund, true
	}
}

func (vm *VendingMachine) Refund(amount int) {
	fmt.Println("Initiating refunding: ", amount)
}
