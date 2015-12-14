package gopath

import (
	"path/filepath"
	"testing"
)

func TestReadmeUsage_withoutGoPath(t *testing.T) {

	var err error
	var p = "/my/path/to/somewhere"

	if p, err = filepath.Abs(p); err != nil {
		// handle error
	}
	if p, err = filepath.EvalSymlinks(p); err != nil {
		// handle error
	}
	if p, err = filepath.Rel(p, "/other/path"); err != nil {
		// handle error
	}

	// go on

}

func TestReadmeUsage_withGoPath(t *testing.T) {
	var otherPath = FromPath("/other/Path")

	var p = FromPath("/my/path/to/somewhere").Abs().EvalSymlinks().Rel(otherPath)

	if p.HasErr() {
		// handle error
	}

	// go on
}

func TestReadmeExtensibility(t *testing.T) {
	var _ = FromPath("/some/path").Do(normalizePath)
}

func normalizePath(p GoPath) GoPath {
	return p.Abs().Clean()
}
