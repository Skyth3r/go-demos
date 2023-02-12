package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please enter only one name")
	} else {
		fmt.Println("Hello " + os.Args[1])
	}
}
