package main

import (
	m "lld/coffeemachine/models"
)

func main() {
	cap := m.NewCoffee("101",
		"Tall Cappuccino",
		10,
		m.Cappuccino,
		m.NewRecipe(
			&m.Ingredients{
				CoffeeBeans: 2,
				Sugar:       1,
				Milk:        2,
				Cream:       1,
			}, "make super simple coffee 101",
		))

	expr := m.NewCoffee("102",
		"Tall Espresso",
		10,
		m.Espresso,
		m.NewRecipe(
			&m.Ingredients{
				CoffeeBeans: 1,
				Sugar:       0,
				Milk:        0,
				Cream:       0,
			}, "make super simple coffee 102",
		))

	cfm := NewCoffeeMachine(map[string]*m.Coffee{cap.ID: cap, expr.ID: expr}, m.NewInventory(10, 10, 10, 10))

	cfm.state.InsertMoney(25)
	cfm.state.Select("101")
	cfm.state.Brew()
}
