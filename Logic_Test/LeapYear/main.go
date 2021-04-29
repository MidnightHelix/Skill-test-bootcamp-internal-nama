package main

import (
	"fmt"
)

func leapYear(start int, end int) []int {

	leapyear := make([]int, 0)
	for i, j := start, 0; i < end; i, j = i+4, j+1 {

		leapyear = append(leapyear, i+4)
	}
	return leapyear
}

func main() {

	var start, end int

	fmt.Println("masukan 2 input tahun")
	fmt.Scanln(&start, &end)

	fmt.Println(leapYear(start, end))

}
