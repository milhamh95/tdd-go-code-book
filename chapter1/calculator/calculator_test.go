package calculator_test

import (
	"testing"

	"tdd-go-code-book/chapter1/calculator"

	"github.com/stretchr/testify/require"
)

func Test_Addition(t *testing.T) {
	res, err := calculator.Addition(1, 2)
	require.Equal(t, 3, res)
	require.NoError(t, err)
}
