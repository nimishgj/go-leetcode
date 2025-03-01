package main

import "math"

func reverse(x int) int {
	negativeNumber := false
	if x < 0 {
		negativeNumber = true
		x *= -1
	}
	reversedX := 0
	for x > 0 {
		reminder := x % 10
		reversedX = (reversedX * 10) + reminder
		x /= 10
	}
	if reversedX > math.MaxInt32 {
		return 0
	}
	if negativeNumber {
		reversedX *= -1
	}
	return reversedX
}
