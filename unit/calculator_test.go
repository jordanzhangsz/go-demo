package unit_test

import (
	"git.webank.io/go_demo/unit"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var calculator = unit.Calculator{}

func TestEvaluatesExpression(t *testing.T) {
	expression := "1+1"

	actual := calculator.Evaluate(expression)

	expected := 2
	if expected != actual {
		t.Errorf("Expected : %v, got : %v", expected, actual)
	}
}

func TestEvaluatesExpression_with_convey(t *testing.T) {
	Convey("Test Evaluates Expression", t, func() {
		Convey("Given 1+1 should return 2", func() {
			actual := calculator.Evaluate("1+1")
			So(actual, ShouldEqual, 2)
		})
	})
}

func TestEvaluatesExpression_with_convey_invalid(t *testing.T) {
	Convey("Test Evaluates Expression", t, func() {
		Convey("Given invalid expression should return error", func() {
			So(func() {calculator.Evaluate("1+abc")}, ShouldPanic)
		})
	})
}
