package calculator_test

import (
	"tdd-go-code-book/chapter2/calculator"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	expectedResult := 2.0
	ce := calculator.Engine{}
	res := ce.Add(1.0, 1.0)
	require.Equal(t, expectedResult, res)
}
