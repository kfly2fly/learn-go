package main

import (
	"fmt"
)

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	hobbies := [3]string{"Coding", "Hiking", "Reading"}
	fmt.Println("My hobbies are:", hobbies)

	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3] // This is some craziness
	fmt.Println(mainHobbies)

	courseGoals := []string{"Learn Go", "Learn all the basics"}
	fmt.Println(courseGoals)

	courseGoals[1] = "Learn all the details"
	newCourseGoals := []string{"Learn all the intricacies"}
	courseGoals = append(courseGoals, newCourseGoals...)
	fmt.Println(courseGoals)

	products := []Product{
		{"My first Product", "first-product", 12.99},
		{"My second Product", "second-product", 129.99},
	}
	fmt.Println(products)
	newProduct := Product{
		"My third Product",
		"third-product",
		15.99,
	}
	products = append(products, newProduct)
	fmt.Println(products)
}

// func main() {
// 	prices := []float64{10.99, 8.99}
// 	fmt.Println(prices[0:1])

// 	updatedPrices := append(prices, 5.99)
// 	fmt.Println(updatedPrices, prices)

// }

// func main() {
// 	var productNames [4]string = [4]string{"A book"}

// 	prices := [4]float64{10, 20.99, 3, 4.0}
// 	fmt.Println(prices)

// 	productNames[2] = "A carpet"
// 	fmt.Println(productNames)

// 	fmt.Println(prices[2])

// 	// A slice does NOT create a new array.
// 	// It holds a reference to a window of the original array
// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 199.99
// 	highlighetPrices := featuredPrices[:1]
// 	fmt.Println(highlighetPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(highlighetPrices), cap(highlighetPrices))

// 	highlighetPrices = highlighetPrices[:3]
// 	fmt.Println(highlighetPrices)
// 	fmt.Println(len(highlighetPrices), cap(highlighetPrices))
// }
