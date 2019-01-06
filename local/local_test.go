/*
Copyright RocKontrol Corp. 2019 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package local

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	tp, _ := ioutil.TempDir(os.TempDir(), "extfs-local-test")
	fs := New(tp)

	f, err := fs.Create("bar/qux")
	require.NoError(t, err)
	assert.Equal(t, filepath.Join(tp, "bar", "qux"), f.Name())
}
