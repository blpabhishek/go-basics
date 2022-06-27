package basic

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "square should square a positive integer",
			args: args{num: 4},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Square(tt.args.num); got != tt.want {
				t.Errorf("Square() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGCD(t *testing.T) {
	actual := GCD(2,3)
	assert.Equal(t, 1, actual, "GCD of 2 and 3 should be 1.")
}

func TestFactorial(t *testing.T) {
	actual := Factorial(5)
	assert.Equal(t, 120, actual)
}

func TestFibonacci(t *testing.T)  {
	fibo:= Fibonacci()
	expected := [5]int{1,2,3,5,8}
	for _, v := range expected {
		assert.Equal(t, v, fibo())
	}
}