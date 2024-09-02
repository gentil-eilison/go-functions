package main

import "fmt"

type transformFunction func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	doubledNumbers := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})
	fmt.Println((doubledNumbers))
}

func transformNumbers(numbers *[]int, transform transformFunction) []int {
	transformedNumbers := []int{}

	for _, number := range *numbers {
		transformedNumbers = append(transformedNumbers, transform(number))
	}
	return transformedNumbers
}