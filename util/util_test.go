/*
Copyright RocKontrol Corp. 2019 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

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

func TestEmptyFilename(t *testing.T) {
	fullpath, err := UnderlyingPath("/cloudchain/test2", "")
	require.NoError(t, err)
	assert.Equal(t, "/cloudchain/test2", fullpath)
}
