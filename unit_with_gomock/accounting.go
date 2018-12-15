package unit_with_gomock

type Person struct {
	SalaryAmount int
	Level        int
}

func CalculateBonus(calculator Calculator, person *Person) int {
	if person.Level > 5 {
		calculator.SetFactor(10)
	}
	return calculator.Calculate(person.SalaryAmount)
}
