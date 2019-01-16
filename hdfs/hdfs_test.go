/*
Copyright RocKontrol Corp. 2019 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package hdfs

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/rkcloudchain/extfs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	hadoopNamenode = "localhost:9000"
)

func TestCreate(t *testing.T) {
	fs, err := New("/cloudchain/test1", &extfs.Config{Addresses: []string{hadoopNamenode}})
	require.NoError(t, err)
	defer fs.Close()

	f, err := fs.Create("myfile.txt")
	require.NoError(t, err)
	defer f.Close()

	_, err = f.Write([]byte("Hello world"))
	require.NoError(t, err)
}

func TestOpen(t *testing.T) {
	fs, err := New("/cloudchain/test1", &extfs.Config{Addresses: []string{hadoopNamenode}})
	require.NoError(t, err)
	defer fs.Close()

	f, err := fs.Open("myfile.txt")
	require.NoError(t, err)
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	require.NoError(t, err)
	assert.Equal(t, "Hello world", string(data))
}

func TestOpenFile(t *testing.T) {
	fs, err := New("/cloudchain/test2", &extfs.Config{Addresses: []string{hadoopNamenode}})
	require.NoError(t, err)
	defer fs.Close()

	f, err := fs.OpenFile("myfile.txt", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	require.NoError(t, err)
	defer f.Close()

	_, err = f.Write([]byte("Hello world"))
	assert.NoError(t, err)
}

func TestAppendFile(t *testing.T) {
	fs, err := New("/cloudchain/test2", &extfs.Config{Addresses: []string{hadoopNamenode}})
	require.NoError(t, err)
	defer fs.Close()

	f, err := fs.OpenFile("myfile.txt", os.O_WRONLY|os.O_APPEND, os.ModePerm)
	require.NoError(t, err)
	defer f.Close()

	_, err = f.Write([]byte(", Xu Qiaolun"))
	assert.NoError(t, err)
}

func TestRemove(t *testing.T) {
	fs, err := New("/cloudchain/test2", &extfs.Config{Addresses: []string{hadoopNamenode}})
	require.NoError(t, err)
	defer fs.Close()

	err = fs.RemoveAll("")
	assert.NoError(t, err)
}

func TestAbs(t *testing.T) {
	fs, err := New("/cloudchain/test2", &extfs.Config{Addresses: []string{hadoopNamenode}})
	require.NoError(t, err)
	defer fs.Close()

	abs, err := fs.Abs("/include")
	require.NoError(t, err)
	assert.Equal(t, "/cloudchain/test2/include", abs)
}
