package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnderlyingPath(t *testing.T) {
	baseDir := "/usr/local"
	filename := "Cellar/go/1.11.4/bin/go"

	fullpath, err := UnderlyingPath(baseDir, filename)

	require.NoError(t, err)
	assert.Equal(t, "/usr/local/Cellar/go/1.11.4/bin/go", fullpath)
}
