package calculator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatorParametrized(t *testing.T) {

	parameters := []struct {
		a        int
		b        int
		expected int
	}{
		{2, 4, 6},
		{1, 8, 9},
		{2, 5, 7},
		{7, 3, 10},
		{22, 4, 26},
		{12, 14, 26},
	}

	for _, p := range parameters {
		t.Run(fmt.Sprintf("%d + %d = %d", p.a, p.b, p.expected), func(t *testing.T) {
			c := Calculator{}
			actual := c.Addition(p.a, p.b)
			assert.Equal(t, p.expected, actual)
		})
	}

}
