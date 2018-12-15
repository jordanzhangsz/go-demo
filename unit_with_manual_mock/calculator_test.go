package unit_with_manual_mock_test

import (
	"git.webank.io/go_demo/unit_with_manual_mock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCalculateBonus_Level_2(t *testing.T) {
	Convey("Test Bonus Calculation", t, func() {
		Convey("Given level 2 person should return double salary as bonus", func() {
			calculator := unit_with_manual_mock.CalculatorDouble{}

			expected := unit_with_manual_mock.CalculateBonus(calculator, &unit_with_manual_mock.Person{100, 2})

			So(expected, ShouldEqual, 200)
		})
	})
}

type mockCalculator struct {
}

func (calculator mockCalculator) Calculate(x int) int {
	return 500
}

func TestCalculateBonus_Level_5(t *testing.T) {
	Convey("Test Bonus Calculation", t, func() {
		Convey("Given level 5 person should return 5 times salary as bonus", func() {
			calculator := mockCalculator{}

			expected := unit_with_manual_mock.CalculateBonus(calculator, &unit_with_manual_mock.Person{100, 5})

			So(expected, ShouldEqual, 500)
		})
	})
}
