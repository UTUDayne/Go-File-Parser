package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var i int = 5

var words = " words words words"

// First comment.

var slicing = []string{" Words ", "words"}

func main() {
	byteString := []byte(fmt.Sprintf("%d", i) + words + "\n")
	file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE, 0644)
	panic_func(err)
	defer file.Close()
	i = i + rand.Intn(10)
	os.WriteFile("test.txt", []byte(""), 0644)
	var word_string = string("")
	for _, value := range slicing {
		word_string += value
	}
	number := 0
	for e := 0; e < 10; e++ {
		number += e
	}
	byteString = append(byteString, append([]byte(strconv.Itoa(number)), []byte(word_string)...)...)
	_, err = file.Write(byteString)
	panic_func(err)
	content, err := os.ReadFile("test.txt")
	string_content := string(content)
	if strings.Contains(string_content, "words") {
		fmt.Println(string_content)
	}
	panic_func(err)
}

func panic_func(err error) {
	if err != nil {
		panic(err)
	}
}
