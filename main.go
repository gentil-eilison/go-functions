package main

import "fmt"

type transformFunction func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(numbers)
	fmt.Println(doubled)
	fmt.Println(tripled)
	fmt.Println("========================")

	transformerFunction1 := getTransformerFunction(&numbers)
	transformerFunction2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformerFunction1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformerFunction2)

	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)
}

func transformNumbers(numbers *[]int, transform transformFunction) []int {
	transformedNumbers := []int{}

	for _, number := range *numbers {
		transformedNumbers = append(transformedNumbers, transform(number))
	}
	return transformedNumbers
}

func getTransformerFunction(numbers *[]int) transformFunction {
	if (*numbers)[0] == 1 {
		return double
	}
	return triple
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}