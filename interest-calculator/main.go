package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// TODO add flags
	balance, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Fatal(err)
	}
	interestRate, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		log.Fatal(err)
	}
	updatedBalance := balance + (balance * (interestRate / 100))

	fmt.Printf("Balance: %.2f\n", updatedBalance)
}
