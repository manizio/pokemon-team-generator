package utils

import "math/rand/v2"

func GenerateRandomNumber() int {
	return rand.IntN(1026 - 1) + 1

}
func GenerateRandomNumberN(max int) int {
	return rand.IntN(max - 1) + 1
}
