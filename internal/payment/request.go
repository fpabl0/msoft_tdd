package payment

// Request defines a payment request.
type Request struct {
	amount float32
}

// NewRequest creates a new payment request.
func NewRequest(amount float32) Request {
	return Request{amount: amount}
}

// GetAmount returns the amount of the request.
func (r *Request) GetAmount() float32 {
	return r.amount
}
