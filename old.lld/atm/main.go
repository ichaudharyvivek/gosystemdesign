package main

import (
	"fmt"
	"lld/atm/models"
	"lld/atm/services"
)

func main() {
	bs := services.NewBankService()
	atm := services.GetATMInstance(bs)

	account, err := bs.CreateAccount("1234567890123456", 10000.0)
	if err != nil {
		fmt.Printf("Error creating account: %v\n", err)
		return
	}

	card, err := bs.CreateCard("0001000200030004", "1234")
	if err != nil {
		fmt.Printf("Error creating card: %v\n", err)
		return
	}

	err = bs.LinkCardToAccount(account, card)
	if err != nil {
		fmt.Printf("Error linking card to account: %v\n", err)
		return
	}

	if err := atm.InsertCard(card); err != nil {
		fmt.Printf("Error inserting card: %v\n", err)
		return
	}

	if err := atm.InsertPin("1234"); err != nil {
		fmt.Printf("Error authenticating card: %v\n", err)
		return
	}

	if err := atm.SelectTransaction(models.TransactionDebit, 2500); err != nil {
		fmt.Printf("Error withdrawing: %v\n", err)
		return
	}

	if err := atm.EjectCard(); err != nil {
		fmt.Printf("Error ejecting card: %v\n", err)
		return
	}
}
