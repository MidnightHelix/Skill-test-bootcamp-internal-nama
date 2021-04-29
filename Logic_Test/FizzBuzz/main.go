package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {

	fizzbuzz := make([]string, 0)

	for i := 1; i <= n; i++ {

		if (i%3 == 0) && (i%5 == 0) {
			fizzbuzz = append(fizzbuzz, "FizzBuzz")
		} else if i%5 == 0 {
			fizzbuzz = append(fizzbuzz, "Buzz")
		} else if i%3 == 0 {
			fizzbuzz = append(fizzbuzz, "Fizz")
		} else {

			fizzbuzz = append(fizzbuzz, strconv.Itoa(i))
		}
	}
	return fizzbuzz

}

func main() {
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		panic("Test")
	}

	fmt.Println(fizzBuzz(i))
}
