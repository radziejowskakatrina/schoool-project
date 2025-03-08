package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(99)
	x := rand.Intn(10)
	y := rand.Intn(10)
	z := x + y
	fmt.Println("Sum of", x, "and", y, "is", z)
}
