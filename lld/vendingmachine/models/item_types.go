package models

type ItemType int

const (
	CHIPS ItemType = iota
	SOFT_DRINKS
	BISCUITS
	SODA
)

func (i ItemType) String() string {
	return [...]string{"Chips", "Soft Drinks", "Biscuits", "Soda"}[i]
}
