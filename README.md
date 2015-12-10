# gopath
An alternative Go API for operating on paths without the need for error handling between each step.

## Usage

Instead of

```go
var err error
var p = "/my/path/to/somewhere"

if p, err = filepath.Abs(p); err != nil {
  // handle error
}
if p, err = filepath.EvalSymlinks(p); err != nil {
  // handle error
}
if p, err = filepath.Rel(p, "other/path"); err != nil {
  // handle error
}

// go on
```

gopath just allows you to

```go
var p = gopath.FromPath("/my/path/to/somewhere").Abs().EvalSymlinks().Rel("other/path")

if p.HasErr() {
  // handle error
}

// go on
```
