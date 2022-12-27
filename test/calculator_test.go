package test

import (
	"testing"
	"unit-testing/utils"

	"github.com/stretchr/testify/assert"
)

func TestCalculator_Multiplying(t *testing.T) {

	type args struct {
		Number1 float64
		Number2 float64
	}

	type want struct {
		response float64
	}

	tests := []struct {
		id   int
		name string
		args args
		want want
	}{
		// Arrange
		{
			id:   1,
			name: "Result should be 36",
			args: args{
				Number1: 6,
				Number2: 6,
			},
			want: want{
				response: 36,
			},
		},
		{
			id:   2,
			name: "Result should be 42",
			args: args{
				Number1: 6,
				Number2: 7,
			},
			want: want{
				response: 42,
			},
		},
		{
			id:   3,
			name: "Result should be 150",
			args: args{
				Number1: 15,
				Number2: 10,
			},
			want: want{
				response: 150,
			},
		},
	}

	for _, tt := range tests {
		currentTest := tt
		t.Run(currentTest.name, func(t *testing.T) {
			t.Parallel()
			var got want

			// Act
			got.response = utils.Mul(currentTest.args.Number1, currentTest.args.Number2)

			// Assert
			assert.Equal(t, currentTest.want.response, got.response, currentTest.name)
		})
	}

}
