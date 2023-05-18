package calculator_test

import (
	"github.com/stretchr/testify/require"
	"tdd-go-code-book/chapter2/calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedResult := 2.0
		ce := calculator.Engine{}
		res := ce.Add(1.0, 1.0)
		require.Equal(t, expectedResult, res)
	})
}
