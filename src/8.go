package main

import "fmt"

func main() {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(10)
	fmt.Println("The generated number is:", number)
}
