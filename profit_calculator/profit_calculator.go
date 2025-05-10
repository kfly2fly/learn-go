package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	revenue, err1 := getUserInput("Revenue: ")
	expenses, err2 := getUserInput("Expenses: ")
	taxRate, err3 := getUserInput("Tax Rate: ")
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.2f\n", ebt)
	fmt.Printf("%.2f\n", profit)
	fmt.Printf("%.2f\n", ratio)
	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.3f", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Value must be a positive number.")
	}

	return userInput, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
