package models

type ItemType int

const (
	CHIPS ItemType = iota
	SODA
	SOFT_DRINKS
	BISCUITS
)

func (r ItemType) String() string {
	return [...]string{"Chips", "Soda", "Soft Drinks", "Biscuits"}[r]
}
