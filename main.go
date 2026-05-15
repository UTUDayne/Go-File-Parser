package main

import (
	"fmt"
	"math/rand"
	"os"
)

var i int = 5

func main() {
	b := []byte(fmt.Sprintf("%d", i))
	err := os.WriteFile("test.txt", b, 0644)
	panic_func(err)
	i = i + rand.Intn(10)
	file, err := os.ReadFile("test.txt")
	fmt.Println(string(file))
}

func panic_func(err error) {
	if err != nil {
		panic(err)
	}
}
