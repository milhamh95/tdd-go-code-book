package calculator_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Addition(t *testing.T) {
	res, err := calculator.Addition(1, 2)
	require.Equal(t, 3, res)
	require.NoError(err)
}
