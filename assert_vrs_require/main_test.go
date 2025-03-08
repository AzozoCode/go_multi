package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func div(a, b int) (float64, error) {
	if a == 0 || b == 0 {
		return 0, fmt.Errorf("numbers must be greater than 0")
	}

	return float64(a) / float64(b), nil
}

type arg struct {
	a int
	b int
}

func Test_Adder(t *testing.T) {
	tests := []struct {
		name   string
		arg    arg
		result float64
	}{
		{"ok", arg{10, 5}, 2},
		{"ok2", arg{6, 4}, 1.5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := div(test.arg.a, test.arg.b)

			if err == nil {
				require.Nil(t, err)
				assert.EqualValues(t, test.result, res)
				assert.EqualValues(t, test.result, res)
			}
		})
	}
}
