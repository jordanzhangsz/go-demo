package unit_with_manual_mock

type CalculatorDouble struct {
}

func (calculatorA CalculatorDouble) Calculate(x int) int {
	return x * 2
}