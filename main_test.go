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
