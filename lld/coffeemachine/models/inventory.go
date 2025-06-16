package models

import "fmt"

type Inventory struct {
	Milk        int
	Cream       int
	Sugar       int
	CoffeeBeans int
}

func NewInventory(milk, cream, sugar, coffeeBeans int) *Inventory {
	return &Inventory{milk, cream, sugar, coffeeBeans}
}

func (i *Inventory) Use(in *Ingredients) bool {
	if i.Milk < 5 || i.Cream < 5 || i.Sugar < 5 || i.CoffeeBeans < 10 {
		fmt.Println("WARNING: Low Inventory. Please refill to continue.")
		return false
	}

	if i.Milk-in.Milk < 0 || i.Cream-in.Cream >= 0 || i.Sugar-in.Sugar >= 0 || i.CoffeeBeans-in.CoffeeBeans >= 0 {
		fmt.Println("Insufficient ingredients...exiting")
		return false
	}

	i.Milk -= in.Milk
	i.Cream -= in.Cream
	i.Sugar -= in.Sugar
	i.CoffeeBeans -= in.CoffeeBeans
	return true
}

func (i *Inventory) Display() {
	fmt.Printf("Inventory - Milk: %dL, Cream: %dmL, Sugar: %dG, CoffeeBeans: %dG\n", i.Milk, i.Cream, i.Sugar, i.CoffeeBeans)
}
