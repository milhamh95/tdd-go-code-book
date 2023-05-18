package format_test

import (
	"errors"
	"github.com/stretchr/testify/require"
	"tdd-go-code-book/chapter2/format"
	"testing"
)

func TestError(t *testing.T) {
	wrappedErr := format.Error("2%3", errors.New("error msg"))

	require.NotNil(t, wrappedErr)
	require.ErrorContains(t, wrappedErr, "error msg")
	require.ErrorContains(t, wrappedErr, "2%3")
}
