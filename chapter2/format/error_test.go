package format

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestError(t *testing.T) {
	wrappedErr := format.Error("2%3", errors.New("error msg"))

	require.NotNil(t, wrappedErr)
	require.Contains(t, wrappedErr, "error msg")
}
