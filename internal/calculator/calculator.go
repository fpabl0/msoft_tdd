package calculator

import "time"

// Calculator defines Calculator object.
type Calculator struct {
	ans int
}

// Addition returns the sum of two numbers.
func (c *Calculator) Addition(a, b int) int {
	return a + b
}

// Subtraction returns the subtraction of two numbers.
func (c *Calculator) Subtraction(a, b int) int {
	return a - b
}

// Division returns the division of two numbers.
func (c *Calculator) Division(a, b int) int {
	return a / b
}

// Multiplication returns the multiplication of two numbers.
func (c *Calculator) Multiplication(a, b int) int {
	time.Sleep(100 * time.Millisecond)
	return a * b
}

// SetAns sets the Ans private variable.
func (c *Calculator) SetAns(v int) { c.ans = v }

// GetAns returns the Ans private variable.
func (c *Calculator) GetAns() int { return c.ans }
