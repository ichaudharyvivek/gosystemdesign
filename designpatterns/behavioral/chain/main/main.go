package main

import "designpatterns/behavioral/chain"

func main() {
	cd := chain.NewCashDispensor()
	cd.Withdraw(2500)
}
