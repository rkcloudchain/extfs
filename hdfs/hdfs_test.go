/*
Copyright RocKontrol Corp. 2019 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package hdfs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	hadoopNamenode = "localhost:9000"
)

func TestCreate(t *testing.T) {
	fs, err := New(hadoopNamenode, "/cloudchain/test1")
	require.NoError(t, err)
	defer fs.Close()

	f, err := fs.Create("myfile.txt")
	require.NoError(t, err)
	defer f.Close()

	_, err = f.Write([]byte("Hello world"))
	require.NoError(t, err)
}
