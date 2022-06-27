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
	assert.Equal(t, actual, 1, "GCD of 2 and 3 should be 1.")
}
