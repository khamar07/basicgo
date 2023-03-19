package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(10)
	fmt.Printf("My favorite number is %d\n", randomNumber)
}
