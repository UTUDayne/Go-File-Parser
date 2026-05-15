package main

import (
	"fmt"
	"math/rand"
)

var i int = 5

func main() {
	i = i + rand.Intn(10)

	fmt.Println("Hello World")
}
