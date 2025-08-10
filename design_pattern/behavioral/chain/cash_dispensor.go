package chain

type CashDispensor struct {
	dispensor Dispensor
}

func NewCashDispensor() *CashDispensor {
	rs100 := NewBaseDispensor(100, 2)
	rs500 := NewBaseDispensor(500, 3)
	rs2000 := NewBaseDispensor(2000, 4)

	rs500.SetNext(rs100)
	rs2000.SetNext(rs500)

	return &CashDispensor{
		dispensor: rs2000,
	}
}

func (c *CashDispensor) Withdraw(amt int) {
	c.dispensor.Dispense(amt)
}
