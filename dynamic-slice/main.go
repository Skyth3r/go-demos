package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	length, _ := strconv.Atoi(os.Args[1])
	slice := make([]int, length)
	for i := 0; i < len(slice); i++ {
		slice[i] = i + 1
	}
	fmt.Printf("%v\n", slice)
}
