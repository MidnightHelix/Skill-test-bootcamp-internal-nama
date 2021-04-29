package main

import (
	"fmt"
)

func FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}

func nearestFibonacci(arr []int) int {

	fibonacci := make([]int, 0)
	var sum, result, near int

	for i := 0; i <= 20; i++ {
		fibonacci = append(fibonacci, FibonacciRecursion(i))
	}

	fmt.Println("Fibonacci sebelum di insert sum dan di sort : ")
	fmt.Println(fibonacci)

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	fibonacci = append(fibonacci, 0)
	copy(fibonacci[1:], fibonacci[0:])
	fibonacci[0] = sum

	for i := 0; i < len(fibonacci)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(fibonacci); j++ {
			if fibonacci[j] < fibonacci[minIndex] {
				fibonacci[j], fibonacci[minIndex] = fibonacci[minIndex], fibonacci[j]
			}
		}
	}

	for i := 0; i < len(fibonacci); i++ {
		if fibonacci[i] == sum {
			if (fibonacci[i+1] - fibonacci[i]) > (fibonacci[i] - fibonacci[i-1]) {
				result = fibonacci[i] - fibonacci[i-1]
				near = fibonacci[i-1]
			} else {
				result = fibonacci[i+1] - fibonacci[i]
				near = fibonacci[i+1]
			}

		}
	}

	fmt.Println("array input : ")
	fmt.Println(arr)
	fmt.Println("sum dari array : ")
	fmt.Println(sum)
	fmt.Println("Fibonacci setelah di insert sum dan di sort : ")
	fmt.Println(fibonacci)
	fmt.Println("nearest fibonacci : ")
	fmt.Println(near)

	fmt.Println("output akhir : ")
	return result
}

func main() {
	arr := make([]int, 0)
	array := [4]int{15, 1, 3, 22} //ubah di sini

	for i := 0; i < 4; i++ { // ubah di sini
		arr = append(arr, array[i])
	}
	fmt.Println(nearestFibonacci(arr))

}
