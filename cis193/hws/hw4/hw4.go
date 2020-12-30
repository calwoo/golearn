package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	FileSum("file_sum.txt", "file_res.txt")
}

// Problem 1a: File processing
// You will be provided an input file consisting of integers, one on each line.
// Your task is to read the input file, sum all the integers, and write the
// result to a separate file.

// FileSum sums the integers in input and writes them to an output file.
// The two parameters, input and output, are the filenames of those files.
// You should expect your input to end with a newline, and the output should
// have a newline after the result.
func FileSum(input, output string) {
	fi, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}

	var res int = 0

	// scan file
	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		res += i
	}

	fo, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}

	fo.WriteString(fmt.Sprint(res))
	defer fo.Close()
}

// Problem 1b: IO processing with interfaces
// You must do the exact same task as above, but instead of being passed 2
// filenames, you are passed 2 interfaces: io.Reader and io.Writer.
// See https://golang.org/pkg/io/ for information about these two interfaces.
// Note that os.Open returns an io.Reader, and os.Create returns an io.Writer.

// IOSum sums the integers in input and writes them to output
// The two parameters, input and output, are interfaces for io.Reader and
// io.Writer. The type signatures for these interfaces is in the Go
// documentation.
// You should expect your input to end with a newline, and the output should
// have a newline after the result.
func IOSum(input io.Reader, output io.Writer) {
	// TODO
}

// Problem 2: Concurrent map access
// Maps in Go [are not safe for concurrent use](https://golang.org/doc/faq#atomic_maps).
// For this assignment, you will be building a custom map type that allows for
// concurrent access to the map using mutexes.
// The map is expected to have concurrent readers but only 1 writer can have
// access to the map.

// PennDirectory is a mapping from PennID number to PennKey (12345678 -> adelq).
// You may only add *private* fields to this struct.
// Hint: Use an embedded sync.RWMutex, see lecture 2 for a review on embedding
type PennDirectory struct {
	directory map[int]string
	mu        sync.RWMutex
}

// Add inserts a new student to the Penn Directory.
// Add should obtain a write lock, and should not allow any concurrent reads or
// writes to the map.
// You may NOT write over existing data - simply raise a warning.
func (d *PennDirectory) Add(id int, name string) {
	_, ok := d.directory[id]
	if !ok {
		d.mu.Lock()
		d.directory[id] = name
		d.mu.Unlock()
	} else {
		fmt.Println("Can't override a value!")
	}
}

// Get fetches a student from the Penn Directory by their PennID.
// Get should obtain a read lock, and should allow concurrent read access but
// not write access.
func (d *PennDirectory) Get(id int) string {
	d.mu.RLock()
	v := d.directory[id]
	d.mu.RUnlock()
	return v
}

// Remove deletes a student to the Penn Directory.
// Remove should obtain a write lock, and should not allow any concurrent reads
// or writes to the map.
func (d *PennDirectory) Remove(id int) {
	d.mu.Lock()
	delete(d.directory, id)
	d.mu.Unlock()
}