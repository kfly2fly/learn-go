package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])
	websites["Facebook"] = "https://facebook.com"

	delete(websites, "Google")
	fmt.Println(websites)

	userNames := make([]string, 2, 5) // slice with len 3 and cap 5
	userNames[0] = "Julie"
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manual")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.6

	courseRatings.output()

	for index, value := range userNames {
		fmt.Println(index)
		fmt.Println(value)
	}

	for key, value := range courseRatings {
		// ..
		fmt.Println(key, value)
	}

}
