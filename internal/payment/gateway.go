package payment

//go:generate mockgen -destination ./mocks_test.go -package payment -mock_names Gateway=GatewayMock . Gateway

// Gateway defines Payment Gateway interface.
type Gateway interface {
	RequestPayment(paymentRequest Request) Response
}
