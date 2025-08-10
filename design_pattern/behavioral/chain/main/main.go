package main

import "design-patterns/behavioral/chain"

func main() {
	cd := chain.NewCashDispensor()
	cd.Withdraw(2500)
}
