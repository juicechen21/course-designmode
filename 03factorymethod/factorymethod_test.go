package factorymethod

import (
	"fmt"
	"testing"
)

func compute(fa OperatorFactory,a,b int) int {
	op := fa.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestOperator(t *testing.T) {
	var fa OperatorFactory

	fa = PlusOperatorFactory{}
	n1 := compute(fa,13,6)


	fa = MinusOperatorFactory{}
	n2 := compute(fa,13,6)

	fmt.Println(n1,n2)
}
