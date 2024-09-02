package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	sum := sumup(numbers...)
	sumTwo := sumup(5, 6, 7, 8)

	fmt.Println(sum)
	fmt.Println(sumTwo)

}

func sumup(numbers ...int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}