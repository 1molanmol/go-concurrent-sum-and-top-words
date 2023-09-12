package cos418_hw1_1

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"sync"
)

// Sum numbers from channel `nums` and output sum to `out`.
// You should only output to `out` once.
func sumWorker(nums chan int, out chan int) {
	s := 0 // Initialize a variable 's' to store the sum of numbers.
	for num := range nums {
		// This loop iterates over values received from the 'nums' channel until it's closed.
		s += num // Add the received 'num' to the running sum 's'.
	}
	out <- s // Send the final sum 's' to the 'out' channel.
}

// Read integers from the file `fileName` and return sum of all values.
// This function must launch `num` go routines running
// `sumWorker` to find the sum of the values concurrently.

// sum calculates the sum of integers read from a file using multiple goroutines.
// It takes two parameters:
//   - num: The number of goroutines (workers) to use for summing.
//   - fileName: The name of the file to read integers from.
//
// It returns the sum of the integers read from the file.
func sum(num int, fileName string) int {
	// Open the file for reading.
	file, err := os.Open(fileName)
	// If there's an error opening the file.
	checkError(err)

	defer file.Close() // Ensure the file is closed when we're done.

	// Read integers from the file.
	nums, err := readInts(file)
	// If there's an error reading integers.
	checkError(err)

	// Create channels for workers and result.
	numsChan := make(chan int, len(nums))
	outChan := make(chan int, num)

	// Create a wait group to wait for all workers to finish.
	var wg sync.WaitGroup

	// Launch `num` goroutines running `sumWorker`.
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sumWorker(numsChan, outChan)
		}()
	}

	// Send numbers to worker channels.
	for _, num := range nums {
		numsChan <- num
	}
	close(numsChan) // Close the input channel when all numbers are sent.

	// Wait for all workers to finish.
	wg.Wait()

	// Close the result channel and collect the results.
	close(outChan)

	sum := 0
	for partialSum := range outChan {
		sum += partialSum
	}

	return sum
}

// Read a list of integers separated by whitespace from `r`.
// Return the integers successfully read with no error, or
// an empty slice of integers and the error that occurred.
func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var elems []int
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return elems, err
		}
		elems = append(elems, val)
	}
	return elems, nil
}
