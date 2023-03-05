package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	inputString := "4242424242424242"
	stringBytes := strings.Split(inputString, "")
	intBytes := make([]int, len(stringBytes))
	checksum := 0

	for index, value := range stringBytes {
		intBytes[index], _ = strconv.Atoi(value)
	}

	for index, value := range intBytes {
		if index%2 == 0 {
			if value*2 > 9 {
				addDigits := 0
				for value > 0 {
					addDigits += value % 10
					value /= 10
				}
				checksum += addDigits
			} else {
				checksum += value * 2
			}
		} else {
			checksum += value
		}
	}

	if checksum%10 == 0 {
		fmt.Print("Valid card number\n")
	} else {
		fmt.Print("Not a valid card number\n")
	}
}
