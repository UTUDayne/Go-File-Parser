package main

import (
	"fmt"
	"math/rand"
	"os"
)

var i int = 5

var words = " words words words"

func main() {
	byteString := []byte(fmt.Sprintf("%d", i) + words + "\n")
	file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE, 0644)
	panic_func(err)
	defer file.Close()
	i = i + rand.Intn(10)
	os.WriteFile("test.txt", []byte(""), 0644)
	_, err = file.Write(byteString)
	panic_func(err)
	content, err := os.ReadFile("test.txt")
	panic_func(err)
	fmt.Println(string(content))
}

func panic_func(err error) {
	if err != nil {
		panic(err)
	}
}
