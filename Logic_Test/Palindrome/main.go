package main

import (
	"fmt"
	"strings"
)

func isPalindrome(str string) string {

	for i, j := 0, len(str)-1; i < len(str); i, j = i+1, j-1 {

		if !strings.EqualFold(string(str[i]), string(str[j])) {
			return "Not Palindrome"
		}
	}
	return "Palindrome"
}

func main() {
	fmt.Println(isPalindrome("Radar"))
	fmt.Println(isPalindrome("Malam"))
	fmt.Println(isPalindrome("Kasur ini rusak"))
	fmt.Println(isPalindrome("ibu Ratna antar ubi"))
	fmt.Println(isPalindrome("Malas"))
	fmt.Println(isPalindrome("Makan nasi goreng"))
	fmt.Println(isPalindrome("Balonku ada lima"))
}
