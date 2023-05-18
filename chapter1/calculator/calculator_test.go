package calculator_test

import (
	"errors"
	"github.com/stretchr/testify/require"
	"tdd-go-code-book/chapter1/calculator"
	"testing"
)

func Test_Addition(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		res, err := calculator.Addition(1, 2)
		require.Equal(t, 3, res)
		require.NoError(t, err)
	})

	t.Run("invalid number", func(t *testing.T) {
		res, err := calculator.Addition(-1, -1)
		require.Equal(t, 0, res)
		require.EqualError(t, errors.New("invalid number"), err.Error())
	})
}
