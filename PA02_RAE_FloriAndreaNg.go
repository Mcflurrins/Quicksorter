// Flori Andrea Ng - 230617173
// PA02_QUICKSORT

package main

// import modules
import (
	"fmt"
	"math/rand"
	"time"
)

// RANDOM SEQUENCE GENERATOR
func randomSequence(length int) []int {
	//make a new slice
	var sequence []int = make([]int, 0)
	max := 1000
	min := -999
	for i := 0; i < length; i++ {
		// produce a random integer in the range [min, max)
		randInt := rand.Intn(max-min) + min
		sequence = append(sequence, randInt)
	}

	//return the slice of random integerss
	return sequence
}

// QUICKSORT ALGORITHM
func Quicksort(seq []int, low int, high int) {
	// stop when the low index is bigger than the high index
	if low < high {
		// find the partition of the current sequence
		part := partition(seq, low, high)

		// recursively call the function on both sides of the partition
		Quicksort(seq, low, part-1)
		Quicksort(seq, part+1, high)
	}
	// we don't need to return anything here
}

// LOMUTO'S PARTITION ALGORITHM
func partition(seq []int, low int, high int) int {
	// take the last element in the sequence as a partition
	pivot := seq[high]
	i := low - 1

	// swap elements at i and j everytime the element at j <= pivot
	for j := low; j < high; j++ {
		if seq[j] <= pivot {
			i += 1
			swap(&seq[i], &seq[j])
		}
	}

	// place the pivot in the correct position by swapping elements at index i and high
	swap(&seq[i+1], &seq[high])

	//return the index of the partition
	return i + 1
}

// SWAP FUNCTION
func swap(element1 *int, element2 *int) {
	*element1, *element2 = *element2, *element1
}

// MAIN FUNCTION
func main() {
	// prompt user for length of the random sequence
	var num int
	fmt.Print("Enter number of elements: ")
	fmt.Scan(&num)

	// create a sequence of random integers
	sequence := randomSequence(num)

	fmt.Println("\nBefore Sorting: \n\n", sequence)

	// get the current time and perform Quicksort
	t := time.Now()
	Quicksort(sequence, 0, len(sequence)-1)

	fmt.Println("\nAfter Sorting: \n\n", sequence, "\n")

	// print the time elapsed to run the program
	fmt.Println("Duration: ", time.Since(t))
}
