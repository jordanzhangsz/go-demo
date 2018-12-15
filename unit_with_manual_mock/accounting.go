package unit_with_manual_mock

type Calculator interface {
	Calculate(x int) int
}

type Person struct {
	SalaryAmount int
	Level        int
}

func CalculateBonus(calculator Calculator, person *Person) int {
	return calculator.Calculate(person.SalaryAmount)
}

