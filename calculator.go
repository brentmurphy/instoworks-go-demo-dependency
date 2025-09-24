package instoworks_go_demo_dependency

import (
	"fmt"
)

func Add(a, b int) int {
	fmt.Printf("calling calc.add [a=%d, b=%d]", a, b)
	return a + b
}
