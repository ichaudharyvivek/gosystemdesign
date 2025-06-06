package models

type Coin int

const (
	ONE  Coin = iota + 1
	TWO       = 2
	FIVE      = 5
	TEN       = 10
)

func (c Coin) String() string {
	switch c {
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
