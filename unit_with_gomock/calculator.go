package unit_with_gomock

type Calculator interface {
	SetFactor(num int)
	Calculate(x int) int
}
