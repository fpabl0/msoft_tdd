package calculator

//go:generate mockgen -destination ./mocks_test.go -package calculator -mock_names ICalculator=CalculatorMock . ICalculator

// ICalculator defines Calculator interface.
type ICalculator interface {
	Addition(a, b int) int
}
