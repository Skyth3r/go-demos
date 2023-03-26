package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	now := time.Now()
	rand.Seed(now.UnixNano())
	sliceSize, _ := strconv.Atoi(os.Args[1])
	randLimit, _ := strconv.Atoi(os.Args[2])
	intSlice := make([]int, sliceSize)

	generateNumbers(intSlice, randLimit)

	fmt.Printf("%v\n", intSlice)
}

func generateNumbers(slice []int, limit int) []int {
	for index := range slice {
		slice[index] = rand.Intn(limit) + 1
	}
	sort.Ints(slice)
	return slice
}
