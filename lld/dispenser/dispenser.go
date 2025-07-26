package dispenser

type Dispenser interface {
	SetNext(next Dispenser)
	Dispense(amount int) error
}
