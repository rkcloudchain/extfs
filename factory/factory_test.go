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

	"github.com/rkcloudchain/extfs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	hadoopNamenode = "localhost:9000"
)

func TestCreateLocalFilesystem(t *testing.T) {
	tp := filepath.Join(os.TempDir(), "extfs-factory-test")
	fs, err := NewFilesystem(fmt.Sprintf("file://%s", tp), &extfs.Config{})
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

func TestCreateHadoopFilesystem(t *testing.T) {
	fs, err := New(fmt.Sprintf("hdfs://%s/opt/hadoop", hadoopNamenode))
	require.NoError(t, err)
	defer fs.Close()

	f, err := fs.OpenFile("hello.txt", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	require.NoError(t, err)
	defer f.Close()

	_, err = f.WriteAt([]byte("hello world"), 10)
	require.EqualError(t, err, "Unsupported operation")

	n, err := f.Write([]byte("hello world"))
	require.NoError(t, err)
	assert.NotZero(t, n)
}
