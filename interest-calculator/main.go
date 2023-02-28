package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {
	years := flag.Float64("y", 1.00, "Number of years interest rate is applied to. Default: 1")
	balance := flag.Float64("b", 100, "Balance. Default: 100")
	inputInterestRate := flag.Float64("r", 1.00, "Interest rate. Default: 1%")
	flag.Parse()

	formattedInterestRate := (*inputInterestRate / 100) + 1
	compoundedInterestRate := math.Pow(formattedInterestRate, *years)
	projectedBalance := *balance * compoundedInterestRate

	fmt.Printf("Initial balance: £%.2f\n", *balance)
	fmt.Printf("Interest rate: %.2f%% \n", *inputInterestRate)
	fmt.Printf("Projected balance: £%.2f\n", projectedBalance)
}
