package models

type Coin int

const (
	ONE  Coin = iota + 1
	TWO       = 2
	FIVE      = 5
	TEN       = 10
)

func (r Coin) String() string {
	switch r {
	case ONE:
		return "One"
	case TWO:
		return "Two"
	case FIVE:
		return "Five"
	case TEN:
		return "Ten"
	default:
		return ""
	}
}
