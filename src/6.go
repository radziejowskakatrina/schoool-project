package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(10)
	fmt.Println("Generated number:", number)
}
