/*
Copyright RocKontrol Corp. 2019 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package factory

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateLocalFilesystem(t *testing.T) {
	tp := filepath.Join(os.TempDir(), "extfs-factory-test")
	fs, err := NewFilesystem(fmt.Sprintf("file://%s", tp))
	require.NoError(t, err)
	defer fs.Close()

	err = fs.MkdirAll("demo1", os.ModePerm)
	require.NoError(t, err)

	f, err := fs.OpenFile("demo1/test.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	require.NoError(t, err)

	_, err = f.Write([]byte("Hello world"))
	require.NoError(t, err)
	err = f.Close()
	require.NoError(t, err)

	f, err = fs.Open("demo1/test.txt")
	require.NoError(t, err)
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	require.NoError(t, err)
	assert.Equal(t, "Hello world", string(data))
}
