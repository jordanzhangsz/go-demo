package unit_with_gomock_test

import (
	"git.webank.io/go_demo/unit_with_gomock"
	"git.webank.io/go_demo/unit_with_gomock/mocks"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCalculateBonus_Level_2(t *testing.T) {
	Convey("Test Bonus Calculation", t, func() {
		Convey("Given level 2 person should return double salary as bonus", func() {
			controller := gomock.NewController(t)
			defer controller.Finish()

			mockCalculator := mock_unit_with_gomock.NewMockCalculator(controller)
			mockCalculator.EXPECT().Calculate(gomock.Any()).Return(200)

			expected := unit_with_gomock.CalculateBonus(mockCalculator, &unit_with_gomock.Person{100, 2})
			So(expected, ShouldEqual, 200)
		})
	})
}

func TestCalculateBonus_Level_5(t *testing.T) {
	Convey("Test Bonus Calculation", t, func() {
		Convey("Given level 5 person should return 5 times salary as bonus", func() {
			controller := gomock.NewController(t)
			defer controller.Finish()

			mockCalculator := mock_unit_with_gomock.NewMockCalculator(controller)
			mockCalculator.EXPECT().Calculate(gomock.Any()).Return(500)

			expected := unit_with_gomock.CalculateBonus(mockCalculator, &unit_with_gomock.Person{100, 5})
			So(expected, ShouldEqual, 500)
		})
	})
}

func TestCalculateBonus_Level_5_Verify_Call_Times(t *testing.T) {
	Convey("Test Bonus Calculation", t, func() {
		Convey("Given level 5 person should not reset the factor of the calculation", func() {
			controller := gomock.NewController(t)
			defer controller.Finish()

			mockCalculator := mock_unit_with_gomock.NewMockCalculator(controller)
			mockCalculator.EXPECT().SetFactor(gomock.Any()).Times(0)
			mockCalculator.EXPECT().Calculate(gomock.Any()).Times(1)
			unit_with_gomock.CalculateBonus(mockCalculator, &unit_with_gomock.Person{100, 5})
		})
	})
}

func TestCalculateBonus_Level_6_Verify_Result(t *testing.T) {
	Convey("Test Bonus Calculation", t, func() {
		Convey("Given level 6 person should return 6 times salary as bonus", func() {
			controller := gomock.NewController(t)
			defer controller.Finish()

			mockCalculator := mock_unit_with_gomock.NewMockCalculator(controller)
			mockCalculator.EXPECT().SetFactor(gomock.Any())
			mockCalculator.EXPECT().Calculate(gomock.Any()).Return(1000)

			expected := unit_with_gomock.CalculateBonus(mockCalculator, &unit_with_gomock.Person{100, 6})
			So(expected, ShouldEqual, 1000)
		})
	})
}

func TestCalculateBonus_Level_6_Verify_Call_Times(t *testing.T) {
	Convey("Test Bonus Calculation", t, func() {
		Convey("Given level 6 person should reset the factor of the calculation", func() {
			controller := gomock.NewController(t)
			defer controller.Finish()

			mockCalculator := mock_unit_with_gomock.NewMockCalculator(controller)
			mockCalculator.EXPECT().SetFactor(gomock.Any()).Times(1)
			mockCalculator.EXPECT().Calculate(gomock.Any()).Times(1)
			unit_with_gomock.CalculateBonus(mockCalculator, &unit_with_gomock.Person{100, 6})
		})
	})
}
