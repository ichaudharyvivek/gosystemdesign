package state

import "fmt"

// Idle state
type IdleState struct {
	Vm *VendingMachine
}

func (s *IdleState) InsertCoin() {
	fmt.Println("Coin inserted.")
	s.Vm.SetState(&HasMoneyState{Vm: s.Vm})
}

func (s *IdleState) SelectProduct() {
	fmt.Println("Insert coin first.")
}

func (s *IdleState) Dispense() {
	fmt.Println("Insert coin and select product first.")
}

// HasMoney state
type HasMoneyState struct {
	Vm *VendingMachine
}

func (s *HasMoneyState) InsertCoin() {
	fmt.Println("Coin already inserted.")
}

func (s *HasMoneyState) SelectProduct() {
	if s.Vm.stock <= 0 {
		fmt.Println("Out of stock.")
		s.Vm.SetState(&OutOfStockState{Vm: s.Vm})
		return
	}
	fmt.Println("Product selected.")
	s.Vm.SetState(&DispensingState{Vm: s.Vm})
}

func (s *HasMoneyState) Dispense() {
	fmt.Println("Select product first.")
}

// Dispensing state
type DispensingState struct {
	Vm *VendingMachine
}

func (s *DispensingState) InsertCoin() {
	fmt.Println("Please wait, dispensing in progress.")
}

func (s *DispensingState) SelectProduct() {
	fmt.Println("Already selected.")
}

func (s *DispensingState) Dispense() {
	fmt.Println("Dispensing product...")
	s.Vm.stock--

	if s.Vm.stock == 0 {
		s.Vm.SetState(&OutOfStockState{Vm: s.Vm})
	} else {
		s.Vm.SetState(&IdleState{Vm: s.Vm})
	}
}

// OutOfStock state
type OutOfStockState struct {
	Vm *VendingMachine
}

func (s *OutOfStockState) InsertCoin() {
	fmt.Println("Out of stock. Coin rejected.")
}

func (s *OutOfStockState) SelectProduct() {
	fmt.Println("Out of stock.")
}

func (s *OutOfStockState) Dispense() {
	fmt.Println("Cannot dispense. Out of stock.")
}
