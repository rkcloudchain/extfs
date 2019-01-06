/*
Copyright RocKontrol Corp. 2019 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package extfs

import "errors"

// errors
var (
	ErrWriteOnly        = errors.New("File can only be written")
	ErrReadOnly         = errors.New("File can only be read")
	ErrUnsupported      = errors.New("Unsupported operation")
	ErrCrossedBoundary  = errors.New("Chroot boundary crossed")
	ErrNeedAbsolutePath = errors.New("We need an absolute path here")
)
