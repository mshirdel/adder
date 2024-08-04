package pkg

import "github.com/mshirdel/adder/internal"

func Client(a, b int) int {
	return internal.Adder(a, b)
}
