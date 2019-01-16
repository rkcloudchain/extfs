A filesystem interface abstraction for Go

[![Build Status](https://travis-ci.org/rkcloudchain/extfs.svg?branch=master)](https://travis-ci.org/rkcloudchain/extfs)

# Overview
extfs is an filesystem interface providing a simple, uniform and universal API interacting with local filesystem or distributed filesystem. extfs has a simple interface design and constructs a file system instance through a factory function. 

The extfs factory method takes a url string as a parameter and creates a corresponding file system instance based on different schemes.

The interface provided by extfs is consistent with the interface provided in the os package.

## Install extfs

First use go get to install the latest version of the library.
```shell
$ go get github.com/rkcloudchain/extfs
```

Next include extfs in your application.
```go
import "github.com/rkcloudchain/extfs/factory"
```

## Declare a filesystem

extfs currently only supports local filesystem and hadoop filesystem.

```go
// local filesystem
fs, err := factory.NewFilesystem("file:///", &extfs.Config{})

or

fs, err := factory.NewFilesystem("hdfs:///", &extfs.Config{User: "hdfsuser"})
```

Then, you can use it like you would the OS package.

## List of all available functions

File System Methods Available:
```go
Create(filename string) (File, error)
Open(filename string) (File, error)
OpenFile(filename string, flag int, perm os.FileMode) (File, error)
Remove(filename string) error
RemoveAll(path string) error
Rename(oldpath, newpath string) error
Stat(filename string) (os.FileInfo, error)
ReadDir(path string) ([]os.FileInfo, error)
MkdirAll(path string, perm os.FileMode) error
Chmod(name string, mode os.FileMode) error
Chtimes(name string, atime time.Time, mtime time.Time) error
Close() error
```
File Interfaces an Methods Available:
```go
io.Closer
io.Reader
io.ReaderAt
io.Seeker
io.Writer
io.WriterAt

Name() string
Stat() (os.FileInfo, error)
Sync() error
Truncate(size int64) error
```

## License
extfs is released under the Apache 2.0 license. See
[LICENSE.txt](https://github.com/rkcloudchain/extfs/blob/master/LICENSE)
