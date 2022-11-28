package main

import (
	"fmt"
)

func main() {
	pwds := make([][]int, 0)

	for i := 11098; i < 98123+1; i++ {
		num := i
		status := false
		fives := make([]int, 0)
		digits := make([]int, 0, 5)

		// Create an array of digits in correct order
		for num != 0 {
			digits = append([]int{num % 10}, digits...)
			num = num / 10
		}

		for idx, d := range digits {
			if d == 5 {
				fives = append(fives, d)
			}

			for i := idx + 1; i < len(digits); i++ {
				if d <= digits[i] {
					status = true
					continue
				}
				status = false
				break
			}

			if !status {
				break
			}
		}

		if status && len(fives) > 1 {
			pwds = append(pwds, digits)
			digits = digits[:0]
		}

		_ = digits
	}

	fmt.Println(len(pwds))
	fmt.Println(pwds[55])
}
