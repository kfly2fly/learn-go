package main

import "fmt"

// func main() {
// 	numbers := []int{1, 2, 3, 4}
// 	moreNumbers := []int{5, 1, 2}

// 	doubled := transformNumbers(&numbers, double)
// 	tripled := transformNumbers(&numbers, triple)

// 	fmt.Println(doubled)
// 	fmt.Println(tripled)

// 	transformerFn1 := getTransformerFunction(&numbers)
// 	transformerFn2 := getTransformerFunction(&moreNumbers)

// 	transformedNumbers := transformNumbers(&numbers, transformerFn1)
// 	moreTransformedNumbers := transformNumbers(&moreNumbers, transformerFn2)

// 	fmt.Println(transformedNumbers)
// 	fmt.Println(moreTransformedNumbers)

// }

type transformFn func(int) int

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFn {
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

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2)
	triple := createTransformer(3)

	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})
	fmt.Println(transformed)

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
