package main

import (
	"errors"
	"math"
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
