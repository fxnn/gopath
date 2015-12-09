package gopath

import "os"

// Path returns the path represented by the given GoPath object.
// If that GoPath is errorneous, a zero value is returned.
func (g GoPath) Path() string {
	return g.path
}

// FileInfo calls Stat() on GoPath and returns the resulting FileInfo.
// When Stat() results in an error, or GoPath is already errorneous,
// the result is a zero interface.
// When the Stat() result is already cached inside the GoPath, Stat() is not
// called again.
func (g GoPath) FileInfo() os.FileInfo {
	if g.fileInfo == nil {
		return g.Stat().FileInfo()
	}
	return g.FileInfo()
}

// FileMode returns g.FileInfo().Mode(), or 0, when FileInfo() fails.
func (g GoPath) FileMode() os.FileMode {
	if info := g.FileInfo(); info != nil {
		return info.Mode()
	}
	return 0
}
