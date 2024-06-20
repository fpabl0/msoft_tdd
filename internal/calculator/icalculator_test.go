package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestICalculator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := NewCalculatorMock(ctrl)
	c.EXPECT().Addition(2, 3).Return(5)

	assert.Equal(t, 5, c.Addition(2, 3))
}
