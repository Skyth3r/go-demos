package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	balance, _ := strconv.Atoi(os.Args[1])
	interestRate, _ := strconv.ParseFloat(os.Args[2], 64)
	newBalance := float64(balance) + (float64(balance) * (interestRate / 100))

	fmt.Printf("Balance: %.2f\n", newBalance)
}
