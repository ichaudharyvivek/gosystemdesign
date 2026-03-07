package main

import (
	"errors"
	"fmt"
	m "lld/coffeemachine/models"
	s "lld/coffeemachine/states"
)

type CoffeeMachine struct {
	balance   int
	selected  *m.Coffee
	coffeeMap map[string]*m.Coffee
	inventory *m.Inventory
	state     s.State
}

func NewCoffeeMachine(cfm map[string]*m.Coffee, inv *m.Inventory) *CoffeeMachine {
	cm := &CoffeeMachine{balance: 0, coffeeMap: cfm, inventory: inv}
	cm.state = s.NewIdleState(cm)
	return cm
}

func (cm *CoffeeMachine) DisplayMenu() {
	fmt.Println("Menu:")
	for k, v := range cm.coffeeMap {
		fmt.Printf("%s: [%s] %s\n", k, v.Type, v.Name)
	}
}

func (cm *CoffeeMachine) DisplayInventory() {
	cm.inventory.Display()
}

// -- INTERFACE METHODS -- //
func (cm *CoffeeMachine) AddBalance(amount int) {
	if amount < 0 {
		fmt.Println("Cannot enter negative amount")
		return
	}

	cm.balance += amount
}

func (cm *CoffeeMachine) GetBalance() int {
	return cm.balance
}

func (cm *CoffeeMachine) SelectCoffee(id string) error {
	found, ok := cm.coffeeMap[id]
	if !ok {
		return errors.New("coffee not found")
	}

	price := found.Price
	if price > cm.balance {
		return errors.New("insufficient balance")
	}

	cm.selected = found
	return nil
}

func (cm *CoffeeMachine) GetSelectedCoffee() *m.Coffee {
	return cm.selected
}

func (cm *CoffeeMachine) SetState(st s.State) {
	cm.state = st
}

func (cm *CoffeeMachine) GetState() s.State {
	return cm.state
}

func (cm *CoffeeMachine) BrewCoffee() (int, error) {
	in := cm.selected.GetIngredients()
	if ok := cm.inventory.Use(in); !ok {
		return 0, errors.New("cannot make coffee")
	}

	refund := max(cm.balance-cm.selected.Price, 0)
	fmt.Printf("Making coffee of type '%s', name:'%s'", cm.selected.Type, cm.selected.Name)
	fmt.Println("Step:", cm.selected.GetInstructions())
	return refund, nil
}
