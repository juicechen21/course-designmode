package builder

import (
	"fmt"
	"testing"
)

func TestBuilder1(t *testing.T) {
	builder := &Builder1{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	fmt.Println(res)
}

func TestBuilder2(t *testing.T) {
	bu := &Builder2{}
	director := NewDirector(bu)
	director.Construct()
	res := bu.GetResult()
	fmt.Println(res)
}
