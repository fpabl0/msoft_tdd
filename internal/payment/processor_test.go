package payment

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type ProcessorSuiteTest struct {
	suite.Suite
	mctrl     *gomock.Controller
	gateway   *GatewayMock
	processor *Processor
}

func (ts *ProcessorSuiteTest) SetupTest() {
	ts.mctrl = gomock.NewController(ts.T())
	ts.gateway = NewGatewayMock(ts.mctrl)
	ts.processor = NewProcessor(ts.gateway)
}

func (ts *ProcessorSuiteTest) Test_given_payment_when_is_correct_then_ok() {
	ts.gateway.EXPECT().RequestPayment(gomock.Any()).Return(NewResponse(StatusOK))
	ts.True(ts.processor.MakePayment(100))
}

func (ts *ProcessorSuiteTest) Test_given_payment_when_is_wrong_then_error() {
	ts.gateway.EXPECT().RequestPayment(gomock.Any()).Return(NewResponse(StatusERROR))
	ts.False(ts.processor.MakePayment(100))
}

func (ts *ProcessorSuiteTest) TearDownTest() {
	ts.mctrl.Finish()
}

func TestProcessorSuiteTest(t *testing.T) {
	suite.Run(t, new(ProcessorSuiteTest))
}
