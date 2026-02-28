package strategy

// PaymentProcessor consolidate the methods
// It is only concerned with paying, it doesn't care what strategy is given to it
type PaymentProcessor struct {
	strategy PaymentStrategy
}

func NewPaymentProcessor(strategy PaymentStrategy) *PaymentProcessor {
	return &PaymentProcessor{strategy}
}

func (p *PaymentProcessor) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentProcessor) PayAmount(amount int) {
	p.strategy.Pay(amount)
}
