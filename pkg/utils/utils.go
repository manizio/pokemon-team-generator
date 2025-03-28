package utils

import "math/rand/v2"

func GenerateRandomNumber() int {
	return rand.IntN(1026 - 1) + 1
}
