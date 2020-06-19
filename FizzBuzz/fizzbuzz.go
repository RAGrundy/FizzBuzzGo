package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(FizzBuzz(i))
	}
}

func FizzBuzz(i int) string {
	if i%5 == 0 && i%3 == 0 {
		return "FizzBuzz"
	} else if i%5 == 0 {
		return "Buzz"
	} else if i%3 == 0 {
		return "Fizz"
	} else {
		return strconv.Itoa(i)
	}
}
