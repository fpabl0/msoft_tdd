package calculator

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type ICalculatorTestSuite struct {
	suite.Suite
	mctrl *gomock.Controller
	calc  *CalculatorMock
}

func (ts *ICalculatorTestSuite) SetupTest() {
	ts.mctrl = gomock.NewController(ts.T())
	ts.calc = NewCalculatorMock(ts.mctrl)
}

func (ts *ICalculatorTestSuite) Test_given_two_integers_when_addition_then_ok() {
	ts.calc.EXPECT().Addition(2, 3).Return(5)
	ts.Equal(5, ts.calc.Addition(2, 3))
	fmt.Println("Test 1")
}

func (ts *ICalculatorTestSuite) Test_given_two_integers_when_subtraction_then_ok() {
	ts.calc.EXPECT().Subtraction(4, 2).Return(2)
	ts.Equal(2, ts.calc.Subtraction(4, 2))
	fmt.Println("Test 2")
}

func (ts *ICalculatorTestSuite) Test_given_two_integers_when_division_then_exception() {
	ts.calc.EXPECT().Division(6, 0).Do(func(a, b int) { panic(divideByZeroPanicValue) })
	ts.PanicsWithValue(divideByZeroPanicValue, func() {
		fmt.Println("Test 3")
		ts.calc.Division(6, 0)
	})
}

func (ts *ICalculatorTestSuite) Test_given_two_integers_when_multiplication_then_timeout() {
	ts.calc.EXPECT().Multiplication(4, 2).DoAndReturn(func(a, b int) int {
		time.Sleep(100 * time.Millisecond)
		return 8
	})
	ts.Eventually(func() bool {
		fmt.Println("Test 4")
		ts.Equal(8, ts.calc.Multiplication(4, 2))
		return true
	}, 150*time.Millisecond, 1*time.Millisecond)
}

func (ts *ICalculatorTestSuite) TearDownTest() {
	ts.mctrl.Finish()
	ts.calc = nil
}

func TestICalculatorTestSuite(t *testing.T) {
	suite.Run(t, new(ICalculatorTestSuite))
}
