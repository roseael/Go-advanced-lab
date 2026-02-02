package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

// Factorial calculates n!
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	if n == 0 {
		return 1, nil
	}
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res, nil
}

// IsPrime checks if a number is prime
func IsPrime(n int) (bool, error) {
	if n < 2 {
		return false, errors.New("prime check requires number >= 2")
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

// Power calculates base^exponent
func Power(base, exponent int) (int, error) {
	if exponent < 0 {
		return 0, errors.New("negative exponents not supported")
	}
	res := 1
	for i := 0; i < exponent; i++ {
		res *= base
	}
	return res, nil
}

// MakeCounter returns a function that increments a private variable
func MakeCounter(start int) func() int {
	count := start
	return func() int {
		count++
		return count
	}
}

// MakeMultiplier captures the factor
func MakeMultiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

// MakeAccumulator returns three functions sharing one state
func MakeAccumulator(initial int) (add func(int), sub func(int), get func() int) {
	acc := initial
	add = func(n int) { acc += n }
	sub = func(n int) { acc -= n }
	get = func() int { return acc }
	return
}

// Apply returns a new slice with the operation applied to each element
func Apply(nums []int, operation func(int) int) []int {
	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = operation(v)
	}
	return result
}

// Filter returns a new slice containing only elements where predicate is true
func Filter(nums []int, predicate func(int) bool) []int {
	var result []int
	for _, v := range nums {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce reduces a slice to a single value
func Reduce(nums []int, initial int, operation func(acc, curr int) int) int {
	accumulator := initial
	for _, v := range nums {
		accumulator = operation(accumulator, v)
	}
	return accumulator
}

// Compose returns a new function f(g(x))
func Compose(f func(int) int, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

// ExploreProcess demonstrates how Go interacts with the OS and process isolation
func ExploreProcess() {
	fmt.Println("====== Process Information ======")

	// A Process ID (PID) is a unique number assigned by the OS to a running program.
	// It's used to manage, track, and signal the program.
	fmt.Printf("Current Process ID: %d\n", os.Getpid())

	// The Parent Process ID (PPID) is the ID of the process that started this one
	// (usually your terminal or IDE).
	fmt.Printf("Parent Process ID: %d\n", os.Getppid())

	data := []int{1, 2, 3, 4, 5}

	// The slice header contains the pointer to the data, length, and capacity.
	fmt.Printf("Memory address of slice header: %p\n", &data)

	// This is the actual memory address where the first integer is stored.
	fmt.Printf("Memory address of first element: %p\n", &data[0])

	/*
	    Process Isolation is a security and stability feature.
	   It ensures that one process cannot read or write to the memory of
	   another process without explicit permission from the OS.
	   This prevents one crashing app from taking down the whole system.
	*/
	fmt.Println("Note: Other processes cannot access these memory addresses due to process isolation")
}

// DoubleValue takes an int. It will NOT modify the original because Go is pass-by-value.
func DoubleValue(x int) {
	x = x * 2
}

// DoublePointer takes a *int. It WILL modify the original because it follows the address.
func DoublePointer(x *int) {
	*x = *x * 2
}

// CreateOnStack returns a value.
func CreateOnStack() int {
	x := 42 // This variable stays on the stack
	return x
}

// CreateOnHeap returns a pointer.
// Because the pointer survives after the function ends, Go moves 'x' to the heap.
func CreateOnHeap() *int {
	x := 42 // This variable escapes to the heap
	return &x
}

func SwapValues(a, b int) (int, int) {
	return b, a
}

func SwapPointers(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

// AnalyzeEscape is for the experiment
func AnalyzeEscape() {
	_ = CreateOnStack()
	_ = CreateOnHeap()
}

/*
ESCAPE ANALYSIS EXPLANATION:
When running 'go build -gcflags "-m"', the variable in CreateOnHeap escapes.
- Which variables escaped? The variable 'x' in CreateOnHeap().
- Why? Because it returns a pointer to a local variable. If it stayed on the stack,
  it would be destroyed when the function returns, making the pointer invalid.
- What does "escapes to heap" mean? It means the memory is allocated in the
  general heap area rather than the function's stack frame, so it can remain
  after the function finishes.
*/
