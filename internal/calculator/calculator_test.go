package calculator

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

var divideByZeroPanicValue = func() (v interface{}) {
	defer func() {
		v = recover()
	}()
	x := 0
	return 1 / x
}()

// Similar to a Java class
type CalculatorTestSuite struct {
	suite.Suite
	calc *Calculator
}

// Equivalent to JUnit @BeforeClass
func (ts *CalculatorTestSuite) SetupSuite() {
	fmt.Println("setup suite")
}

// Equivalent to JUnit @Before
func (ts *CalculatorTestSuite) SetupTest() {
	fmt.Println("setup")
	ts.calc = &Calculator{}
}

// Prefix "Test" in function name is equivalent to JUnit @Test
func (ts *CalculatorTestSuite) Test_given_two_integers_when_addition_then_ok() {
	ts.Equal(6, ts.calc.Addition(4, 2))
	fmt.Println("Test 1")
}

func (ts *CalculatorTestSuite) Test_given_two_integers_when_subtraction_then_ok() {
	ts.Equal(2, ts.calc.Subtraction(4, 2))
	fmt.Println("Test 2")
}

func (ts *CalculatorTestSuite) Test_given_two_integers_when_division_then_exception() {
	ts.PanicsWithValue(divideByZeroPanicValue, func() {
		fmt.Println("Test 3")
		ts.calc.Division(6, 0)
	})
}

func (ts *CalculatorTestSuite) Test_given_two_integers_when_multiplication_then_timeout() {
	ts.Eventually(func() bool {
		fmt.Println("Test 4")
		ts.Equal(8, ts.calc.Multiplication(4, 2))
		return true
	}, 150*time.Millisecond, 1*time.Millisecond)
}

// Equivalent to JUnit @After
func (ts *CalculatorTestSuite) TearDownTest() {
	fmt.Println("teardown")
	ts.calc.SetAns(0)
	ts.calc = nil
}

// Equivalent to JUnit @AfterClass
func (ts *CalculatorTestSuite) TearDownSuite() {
	fmt.Println("teardown suite")
}

// Needed to actually run the suite tests.
func TestCalculatorTestSuite(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))
}
