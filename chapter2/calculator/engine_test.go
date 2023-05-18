package calculator_test

import (
	"github.com/stretchr/testify/require"
	"tdd-go-code-book/chapter2/calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	ce := calculator.Engine{}

	actAssert := func(x, y, want float64) {
		// act
		got := ce.Add(x, y)

		// Assert
		require.Equal(t, want, got)
	}

	t.Run("success", func(t *testing.T) {
		// Arrange
		expectedResult := 2.0

		// Act
		res := ce.Add(1.0, 1.0)

		// Assert
		require.Equal(t, expectedResult, res)
	})

	t.Run("negative number", func(t *testing.T) {
		x := -1.0
		y := -2.0

		actAssert(x, y, -3.0)
	})
}
