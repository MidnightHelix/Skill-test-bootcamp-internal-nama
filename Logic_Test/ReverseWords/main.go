package main

import (
	"fmt"
	"strings"
)

var result string

func main() {
	str := "I am A Great human"
	reverse := []string{}
	strArray := strings.Fields(str)
	var satuan string
	for _, f := range strArray {
		abjad := f
		for _, c := range abjad {
			result = string(c) + result
		}
		satuan = result
		result = ""
		reverse = append(reverse, satuan)
	}
	//convert slice to string///
	str2 := strings.Join(reverse, " ")
	fmt.Println(str2)
}
