package main

import "fmt"

func main() {
	t := make(chan int, 5)
	t <- 2
	t <- 3
	t <- -1
	t <- 4
	d := Filter(t, func(n int) bool { return n > 0 })
	for v := range d {
		fmt.Println(v)
	}
}

// Filter copies values from the input channel into an output channel that match the filter function p
// The function p determines whether an int from the input channel c is sent on the output channel
func Filter(c <-chan int, p func(int) bool) <-chan int {
	chanLen := len(c)
	d := make(chan int, chanLen)
	for i := 0; i < chanLen; i++ {
		v := <-c
		if p(v) {
			d <- v
		}
	}
	return d
}

// Result is a type representing a single result with its index from a slice
type Result struct {
	index  int
	result string
}

// ConcurrentRetry runs all the tasks concurrently and sends the output in a Result channel
//
// concurrent is the limit on the number of tasks running in parallel. Your
// solution must not run more than `concurrent` number of tasks in parallel.
//
// retry is the number of times that the task should be attempted. If a task
// returns an error, the function should be retried immediately up to `retry`
// times. Only send the results of a task into the output channel if it does not error.
//
// Multiple instances of ConcurrentRetry should be able to run simultaneously
// without interfering with one another, so global variables should not be used.
// The function must return the channel without waiting for the tasks to
// execute, and all results should be sent on the output channel. Once all tasks
// have been completed, close the channel.
func ConcurrentRetry(tasks []func() (string, error), concurrent int, retry int) <-chan Result {
	runningTasks := make(chan int, concurrent)
	r := make(chan string, concurrent)
	rs := make(chan Result, concurrent)
	for i, task := range tasks {
		runningTasks <- i
		go func() {
			for i := 0; i < retry; i++ {
				s, err := task()
				if err == nil {
					r <- s
				}
			}
		}()
	}

	ind := 0
	for s := range r {
		rs <- Result{index: ind, result: s}
		ind++
	}

	return rs
}

// Task is an interface for types that process integers
type Task interface {
	Execute(int) (int, error)
}

// Fastest returns the result of the fastest running task
// Fastest accepts any number of Task structs. If no tasks are submitted to
// Fastest(), it should return an error.
// You should return the result of a Task even if it errors.
// Do not leave any pending goroutines. Make sure all goroutines are cleaned up
// properly and any synchronizing mechanisms closed.
func Fastest(input int, tasks ...Task) (int, error) {
	// TODO
	return 0, nil
}

// MapReduce takes any number of tasks, and feeds their results through reduce
// If no tasks are supplied, return an error.
// If any of the tasks error during their execution, return an error immediately.
// Once all tasks have completed successfully, return the value of reduce on
// their results in any order.
// Do not leave any pending goroutines. Make sure all goroutines are cleaned up
// properly and any synchronizing mechanisms closed.
func MapReduce(input int, reduce func(results []int) int, tasks ...Task) (int, error) {
	// TODO
	return 0, nil
}
