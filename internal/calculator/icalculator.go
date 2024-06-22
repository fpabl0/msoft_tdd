package calculator

//go:generate mockgen -destination ./mocks_test.go -package calculator -mock_names ICalculator=CalculatorMock . ICalculator

// ICalculator defines Calculator interface that represents a web service.
type ICalculator interface {
	Addition(a, b int) int
	Subtraction(a, b int) int
	Multiplication(a, b int) int
	Division(a, b int) int
}
