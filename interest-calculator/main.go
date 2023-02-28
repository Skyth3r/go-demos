package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	balance, _ := strconv.ParseFloat(os.Args[1], 64)
	interestRate, _ := strconv.ParseFloat(os.Args[2], 64)
	newBalance := balance + (balance * (interestRate / 100))

	fmt.Printf("Balance: %.2f\n", newBalance)
}
