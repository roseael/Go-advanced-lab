package main

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    int
		wantErr bool
	}{
		{"0!", 0, 1, false},
		{"5!", 5, 120, false},
		{"3!", 3, 6, false},
		{"Negative", -1, 0, true},
		{"1!", 1, 1, false},
		{"Large-ish", 10, 3628800, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Factorial(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Factorial() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    bool
		wantErr bool
	}{
		{"2 is prime", 2, true, false},
		{"3 is prime", 3, true, false},
		{"4 is composite", 4, false, false},
		{"17 is prime", 17, true, false},
		{"25 is composite", 25, false, false},
		{"Edge case 1", 1, false, true},
		{"Edge case 0", 0, false, true},
		{"Negative number", -5, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsPrime(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsPrime(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsPrime(%d) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		name     string
		base     int
		exponent int
		want     int
		wantErr  bool
	}{
		{"2 to the 3", 2, 3, 8, false},
		{"5 to the 0", 5, 0, 1, false},
		{"0 to the 5", 0, 5, 0, false},
		{"10 to the 1", 10, 1, 10, false},
		{"Negative exponent", 2, -1, 0, true},
		{"3 squared", 3, 2, 9, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Power(tt.base, tt.exponent)
			if (err != nil) != tt.wantErr {
				t.Errorf("Power(%d, %d) error = %v, wantErr %v", tt.base, tt.exponent, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Power(%d, %d) = %v, want %v", tt.base, tt.exponent, got, tt.want)
			}
		})
	}
}

func TestMakeCounter(t *testing.T) {
	t.Run("Counter increments correctly", func(t *testing.T) {
		counter := MakeCounter(0)
		if got := counter(); got != 1 {
			t.Errorf("First call = %d, want 1", got)
		}
		if got := counter(); got != 2 {
			t.Errorf("Second call = %d, want 2", got)
		}
	})

	t.Run("Counters are independent", func(t *testing.T) {
		c1 := MakeCounter(0)
		c2 := MakeCounter(10)

		c1() // becomes 1
		c2() // becomes 11

		if got1, got2 := c1(), c2(); got1 != 2 || got2 != 12 {
			t.Errorf("Counters interfered: c1=%d, c2=%d", got1, got2)
		}
	})
}

func TestMakeMultiplier(t *testing.T) {
	tests := []struct {
		name   string
		factor int
		input  int
		want   int
	}{
		{"Double 5", 2, 5, 10},
		{"Triple 10", 3, 10, 30},
		{"Multiply by 0", 0, 100, 0},
		{"Multiply by negative", -2, 4, -8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			multiplier := MakeMultiplier(tt.factor)
			if got := multiplier(tt.input); got != tt.want {
				t.Errorf("Multiplier(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestMakeAccumulator(t *testing.T) {
	add, sub, get := MakeAccumulator(100)

	t.Run("Add and Get", func(t *testing.T) {
		add(50)
		if got := get(); got != 150 {
			t.Errorf("After adding 50, got %d, want 150", got)
		}
	})

	t.Run("Subtract and Get", func(t *testing.T) {
		sub(30)
		if got := get(); got != 120 {
			t.Errorf("After subtracting 30, got %d, want 120", got)
		}
	})

	t.Run("Independent state", func(t *testing.T) {
		_, _, get2 := MakeAccumulator(0)
		if got1, got2 := get(), get2(); got1 == got2 {
			t.Errorf("Accumulators shared state! Acc1: %d, Acc2: %d", got1, got2)
		}
	})
}

func TestHigherOrder(t *testing.T) {
	nums := []int{1, 2, 3, 4}

	t.Run("Apply: Square numbers", func(t *testing.T) {
		square := func(x int) int { return x * x }
		got := Apply(nums, square)
		want := []int{1, 4, 9, 16}
		for i := range got {
			if got[i] != want[i] {
				t.Errorf("Apply failed, got %v, want %v", got, want)
			}
		}
	})

	t.Run("Filter: Even numbers", func(t *testing.T) {
		isEven := func(x int) bool { return x%2 == 0 }
		got := Filter(nums, isEven)
		want := []int{2, 4}
		if len(got) != len(want) || got[0] != 2 || got[1] != 4 {
			t.Errorf("Filter failed, got %v, want %v", got, want)
		}
	})

	t.Run("Reduce: Sum", func(t *testing.T) {
		sum := func(acc, curr int) int { return acc + curr }
		got := Reduce(nums, 0, sum)
		if got != 10 {
			t.Errorf("Reduce sum failed, got %d, want 10", got)
		}
	})

	t.Run("Compose: Double then add two", func(t *testing.T) {
		addTwo := func(x int) int { return x + 2 }
		double := func(x int) int { return x * 2 }
		combined := Compose(addTwo, double) // f(g(x)) -> (5*2) + 2
		if got := combined(5); got != 12 {
			t.Errorf("Compose failed, got %d, want 12", got)
		}
	})
}

func TestPointers(t *testing.T) {
	t.Run("SwapValues should not affect originals", func(t *testing.T) {
		a, b := 5, 10
		SwapValues(a, b)
		if a != 5 || b != 10 {
			t.Errorf("SwapValues modified originals: a=%d, b=%d", a, b)
		}
	})

	t.Run("SwapPointers should affect originals", func(t *testing.T) {
		a, b := 5, 10
		SwapPointers(&a, &b)
		if a != 10 || b != 5 {
			t.Errorf("SwapPointers failed: a=%d, b=%d", a, b)
		}
	})
}
