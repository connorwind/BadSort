package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	slice := generateslice(10)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	numbersort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
	fmt.Println("\n--- Time Taken \n\n", time.Since(start),"\n")
}

// Generates a slice of size 20, size filled with random numbers 0-999999
func generateslice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(20)
	}
	return slice
}
//sorts the list by placing in a new list based on number size
func numbersort(unsort []int) {
	var n = len(unsort)
	//Declare new slice that will be sorted
	sort := make([]int,n,n)
	
	//count to determine that all values have been found from unsort
	var count = 0
	//loop to sort
	for i := 0; count<n; i++{
		for j :=0; j<n; j++{
			if unsort[j] == i {
				sort[count] = unsort[j]
				count++
			} 
		}
	}
//Puts sorted values from new slice back into original slice
	for i := 0; i<n; i++{
		unsort[i] = sort[i]
	}
}