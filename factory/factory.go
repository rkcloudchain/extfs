/*
Copyright RocKontrol Corp. 2019 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package factory

import (
	"errors"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/rkcloudchain/extfs"
	"github.com/rkcloudchain/extfs/hdfs"
	"github.com/rkcloudchain/extfs/local"
)

const (
	defaultURL = "file:///"
)

// NewFilesystem returns a filesystem based on url
func NewFilesystem(u string) (extfs.Filesystem, error) {
	url, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	scheme := url.Scheme
	authority := url.Host
	path := url.Path

	if scheme == "" && authority == "" && path == "" {
		return NewFilesystem(defaultURL)
	}
	if scheme != "" && authority == "" && path == "" {
		defaultURI, _ := url.Parse(defaultURL)
		if defaultURI.Scheme == scheme {
			return NewFilesystem(defaultURL)
		}
	}

	return createFileSystem(url)
}

func createFileSystem(url *url.URL) (extfs.Filesystem, error) {
	lower := strings.ToLower(url.Scheme)
	switch lower {
	case "file":
		base, err := getBaseDir(url)
		if err != nil {
			return nil, err
		}
		return local.New(base), nil

	case "hdfs":
		authority := url.Host
		if authority == "" {
			return nil, errors.New("The host can't be empty")
		}
		base, err := getBaseDir(url)
		if err != nil {
			return nil, err
		}
		return hdfs.New(authority, base)

	default:
		return nil, fmt.Errorf("Unsupported filesystem %s", lower)
	}
}

func getBaseDir(url *url.URL) (string, error) {
	base := url.Path
	if base == "" {
		base = "/"
	}

	if !filepath.IsAbs(base) {
		return "", extfs.ErrNeedAbsolutePath
	}

	return base, nil
}
