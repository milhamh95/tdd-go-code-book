package format_test

import (
	"errors"
	"github.com/stretchr/testify/require"
	"tdd-go-code-book/chapter2/format"
	"testing"
)

func TestError(t *testing.T) {
	expectedErr := errors.New("error msg")
	expr := "2%3"
	wrappedErr := format.Error(expr, expectedErr)

	require.NotNil(t, wrappedErr)
	require.ErrorContains(t, wrappedErr, expectedErr.Error())
	require.ErrorContains(t, wrappedErr, expr)
}
