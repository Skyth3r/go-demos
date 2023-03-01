package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	inputString := "4003600000000014"
	stringBytes := strings.Split(inputString, "")
	intBytes := make([]int, len(stringBytes))
	checksum := 0

	for index, value := range stringBytes {
		intBytes[index], _ = strconv.Atoi(value)
	}

	for index, value := range intBytes {
		if index%2 != 0 {
			// TODO add a check for if value * 2 is larger than 9
			if value*2 > 9 {
				// TODO split number and add the digits together
			} else {
				checksum += value * 2
			}
		} else {
			checksum += value
		}
	}

	fmt.Printf("Checksum: %d\n", checksum)

	if checksum%10 != 0 {
		fmt.Print("Not a valid card number\n")
	} else {
		fmt.Print("Valid card number\n")
	}
}
