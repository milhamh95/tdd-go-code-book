package calculator_test

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAdd(t *testing.T) {
	res := calculator.Add(1, 1)

	require.Equal(t, 2, res)
}
