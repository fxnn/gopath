package gopath

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func testComponents(t *testing.T) {

	assertStrSliceEqual(t, FromPath("").Components(), "")
	assertStrSliceEqual(t, FromPath(string(os.PathSeparator)).Components(), "")
	assertStrSliceEqual(t, FromPath("a").Components(), "a")
	assertStrSliceEqual(t, FromPath(filepath.Join("a", "b", "c")).Components(), "a", "b", "c")

}

func testComponents_errorneous(t *testing.T) {

	assertSliceEmpty(t, FromErr(errors.New("test err")).Components())

}
