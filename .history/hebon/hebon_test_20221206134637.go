package hebon

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHebonConvert(t *testing.T) {
	hebon :=ToHebon("こんにちは")
	require.Equal(t, hebon, "konnichiha")
}

