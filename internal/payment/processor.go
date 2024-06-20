package payment

// Processor defines processor object.
type Processor struct {
	gateway Gateway
}

// NewProcessor creates a new processor.
func NewProcessor(gateway Gateway) *Processor {
	return &Processor{gateway: gateway}
}

// MakePayment performs a payment using the internal gateway.
func (p *Processor) MakePayment(amount float32) bool {
	response := p.gateway.RequestPayment(NewRequest(amount))
	if response.GetStatus() == StatusOK {
		return true
	}
	return false
}
