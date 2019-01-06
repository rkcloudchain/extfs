/*
Copyright RocKontrol Corp. 2019 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package hdfs

import (
	"os"

	"github.com/colinmarc/hdfs/v2"
	"github.com/rkcloudchain/extfs"
)

type file struct {
	reader *hdfs.FileReader
	writer *hdfs.FileWriter
}

func newFile(reader *hdfs.FileReader, writer *hdfs.FileWriter) *file {
	return &file{reader: reader, writer: writer}
}

func (f *file) Close() error {
	if f.reader != nil {
		return f.reader.Close()
	} else if f.writer != nil {
		return f.writer.Close()
	} else {
		return nil
	}
}

func (f *file) Read(p []byte) (int, error) {
	if f.reader == nil {
		return 0, extfs.ErrWriteOnly
	}

	return f.reader.Read(p)
}

func (f *file) ReadAt(p []byte, off int64) (int, error) {
	if f.reader == nil {
		return 0, extfs.ErrWriteOnly
	}

	return f.reader.ReadAt(p, off)
}

func (f *file) Seek(offset int64, whence int) (int64, error) {
	if f.reader == nil {
		return 0, extfs.ErrUnsupported
	}

	return f.reader.Seek(offset, whence)
}

func (f *file) Write(p []byte) (int, error) {
	if f.writer == nil {
		return 0, extfs.ErrReadOnly
	}

	return f.writer.Write(p)
}

func (f *file) WriteAt(p []byte, off int64) (int, error) {
	return 0, extfs.ErrUnsupported
}

func (f *file) Name() string {
	if f.reader != nil {
		return f.reader.Name()
	}
	return ""
}

func (f *file) Stat() (os.FileInfo, error) {
	if f.reader != nil {
		return f.reader.Stat(), nil
	}

	return nil, extfs.ErrUnsupported
}

func (f *file) Sync() error {
	if f.writer != nil {
		return f.writer.Flush()
	}

	return nil
}

func (f *file) Truncate(size int64) error {
	return extfs.ErrUnsupported
}
