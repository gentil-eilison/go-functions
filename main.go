package main

import "fmt"

type transformFunction func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	double := createTransformer(2)
	triple := createTransformer(3)
	doubledNumbers := transformNumbers(&numbers, double)
	tripledNumbers := transformNumbers(&numbers, triple)

	fmt.Println(numbers)
	fmt.Println(doubledNumbers)
	fmt.Println(tripledNumbers)
}

func transformNumbers(numbers *[]int, transform transformFunction) []int {
	transformedNumbers := []int{}

	for _, number := range *numbers {
		transformedNumbers = append(transformedNumbers, transform(number))
	}
	return transformedNumbers
}

func createTransformer(factor int) transformFunction {
	return func(number int) int {
		return number * factor
	}
}