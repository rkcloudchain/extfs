package util

import (
	"path/filepath"
	"strings"

	"github.com/rkcloudchain/extfs"
)

// UnderlyingPath returns the full path of the merged filename with basedir
func UnderlyingPath(baseDir, filename string) (string, error) {
	if isCrossBoundaries(filename) {
		return "", extfs.ErrCrossedBoundary
	}

	return filepath.Join(baseDir, filename), nil
}

func isCrossBoundaries(path string) bool {
	path = filepath.ToSlash(path)
	path = filepath.Clean(path)

	return strings.HasPrefix(path, ".."+string(filepath.Separator))
}
