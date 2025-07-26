package models

type Card struct {
	cardNumber string
	pin        string
}

func NewCard(cardNumber, pin string) *Card {
	return &Card{
		cardNumber: cardNumber,
		pin:        pin,
	}
}

func (c *Card) GetCardNumber() string {
	return c.cardNumber
}

func (c *Card) ValidatePin(pin string) bool {
	return c.pin == pin
}
