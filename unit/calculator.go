package unit

import (
	"fmt"
	"strconv"
	"strings"
)

type Calculator struct {
}

func (calculator *Calculator) Evaluate(expression string) int {
	sum := 0
	slice := strings.Split(expression, "+")
	for i := 0; i < len(slice); i++ {
		num, err := strconv.Atoi(slice[i])
		if err != nil {
			panic(fmt.Errorf("The expression is invalid - %v", expression))
		}
		sum = sum + num
	}
	return sum
}
