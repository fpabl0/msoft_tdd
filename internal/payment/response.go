package payment

// Status represents the payment status.
type Status int

// Payment status options
const (
	StatusOK Status = iota
	StatusERROR
)

// Response defines payment response.
type Response struct {
	status Status
}

// NewResponse creates a new payment response.
func NewResponse(s Status) Response {
	return Response{status: s}
}

// GetStatus gets the payment response status.
func (r *Response) GetStatus() Status {
	return r.status
}
