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

// New returns a filesystem based on url and options
func New(u string, opts ...extfs.ClientOption) (extfs.Filesystem, error) {
	cfg := &extfs.Config{}
	for _, option := range opts {
		err := option(cfg)
		if err != nil {
			return nil, errors.New("Failed to read opts")
		}
	}
	return NewFilesystem(u, cfg)
}

// NewFilesystem returns a filesystem based on url
func NewFilesystem(u string, cfg *extfs.Config) (extfs.Filesystem, error) {
	url, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	scheme := url.Scheme
	authority := url.Host
	path := url.Path

	if scheme == "" && authority == "" && path == "" {
		return NewFilesystem(defaultURL, cfg)
	}
	if scheme != "" && authority == "" && path == "" {
		defaultURI, _ := url.Parse(defaultURL)
		if defaultURI.Scheme == scheme {
			return NewFilesystem(defaultURL, cfg)
		}
	}

	return createFileSystem(url, cfg)
}

func createFileSystem(url *url.URL, cfg *extfs.Config) (extfs.Filesystem, error) {
	lower := strings.ToLower(url.Scheme)
	switch lower {
	case "file":
		base, err := getBaseDir(url)
		if err != nil {
			return nil, err
		}
		return local.New(base), nil

	case "hdfs":
		if cfg == nil {
			cfg = &extfs.Config{}
		}

		authority := url.Host
		if authority != "" {
			cfg.Addresses = append(cfg.Addresses, authority)
		}
		base, err := getBaseDir(url)
		if err != nil {
			return nil, err
		}

		return hdfs.New(base, cfg)

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
