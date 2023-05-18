package format_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"tdd-go-code-book/chapter2/format"
	"testing"
)

func TestResult(t *testing.T) {
	// Arrange
	result := 5.55
	expr := "2%3"

	// Act
	wrappedRes := format.Result(expr, result)

	// Assert.
	require.Contains(t, wrappedRes, expr)
	require.Contains(t, wrappedRes, fmt.Sprint(result))
}
