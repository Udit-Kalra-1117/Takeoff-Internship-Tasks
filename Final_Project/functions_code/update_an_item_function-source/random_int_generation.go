package helloworld

import "math/rand"

// randomInt generates a random number to be used as a query parameter in image URLs
func random_int() int {
	return rand.Intn(9999999999-1000000000+1) + 1000000000
}