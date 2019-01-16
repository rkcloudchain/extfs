/*
Copyright RocKontrol Corp. 2019 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package local

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/rkcloudchain/extfs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	tp := filepath.Join(os.TempDir(), "extfs-local-test")
	fs := New(tp)

	f, err := fs.Create("bar/qux")
	require.NoError(t, err)
	defer f.Close()
	assert.Equal(t, filepath.Join(tp, "bar", "qux"), f.Name())
}

func TestCreateErrCrossedBoundary(t *testing.T) {
	fs := New("/foo")
	_, err := fs.Create("../foo")
	assert.Equal(t, extfs.ErrCrossedBoundary, err)
}

func TestOpen(t *testing.T) {
	tp := filepath.Join(os.TempDir(), "extfs-local-test")
	fs := New(tp)

	f, err := fs.Open("bar/qux")
	require.NoError(t, err)
	defer f.Close()
	assert.Equal(t, filepath.Join(tp, "bar", "qux"), f.Name())
}

func TestOpenFile(t *testing.T) {
	tp := filepath.Join(os.TempDir(), "extfs-local-test")
	fs := New(tp)

	f, err := fs.OpenFile("bar/qux", os.O_RDONLY, 0666)
	require.NoError(t, err)
	defer f.Close()
	assert.Equal(t, filepath.Join(tp, "bar", "qux"), f.Name())
}

func TestRemove(t *testing.T) {
	tp := filepath.Join(os.TempDir(), "extfs-local-test")
	fs := New(tp)

	err := fs.Remove("bar/qux")
	require.NoError(t, err)

	_, err = fs.Stat("bar/qux")
	assert.True(t, os.IsNotExist(err))
}
