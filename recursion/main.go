package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumup(numbers)

	fmt.Println(sum)

	sum = sumup2(1, 2, 3, 4, 5, 6)
	fmt.Println(sum)

	anotherSum := sumup2(1, numbers...)
	fmt.Println(anotherSum)
}

func factorial(number int) int {
	// result := 1

	// for i := 1; i <= number; i++ {
	// 	result = result * i
	// }

	// return result

	if number == 0 {
		return 1
	}

	return number * factorial(number-1)
}

func sumup(numbers []int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}

func sumup2(startingValue int, numbers ...int) int {
	sum := startingValue

	for _, val := range numbers {
		sum += val
	}

	return sum
}
