package instoworks_go_demo_dependency

import (
	"fmt"
)

func Add(a, b int) int {
	fmt.Printf("calling add it [a=%d, b=%d]", a, b)
	return a + b
}
