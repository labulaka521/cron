package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println((1<<10) & (1 << 11 | 1 << 10))
	// 10  100 1000
	// 1110
	fmt.Println(1<<1 | 1<<2 | 1<<3)
	// 10, 100, 1000
	fmt.Println(1<<1, 1<<2, 1<<3)

	// 1110  0010
	// 10
	fmt.Println(14 & (1 << 1))
	// 1110  0100

	fmt.Println(GetNumber(1<<1 | 1<<5 | 1<<10))

	var a uint64 = ^(math.MaxUint64 << (60 + 1)) & (math.MaxUint64 << 1)
	fmt.Println(a)


	
}

// GetNumber return all bits
func GetNumber(n int) []int {
	res := []int{}

	for i := 2; i < n; i = i * 2 {
		if i&n == 0 {
			continue
		}
		res = append(res, int(math.Log2(float64(i))))
	}
	return res
}
