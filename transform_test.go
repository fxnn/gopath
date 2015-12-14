package gopath

import "testing"

func TestRel_bc(t *testing.T) {

	var base = FromPath("/a")
	var target = FromPath("/b/c")

	var rel = base.Rel(target)
	assertGoPathEqual(t, rel, FromPath("../b/c"))

	var relTo = target.RelTo(base)
	assertGoPathEqual(t, relTo, FromPath("../b/c"))

}

func TestRel_abc(t *testing.T) {

	var base = FromPath("/a")
	var target = FromPath("/a/b/c")

	var rel = base.Rel(target)
	assertGoPathEqual(t, rel, FromPath("b/c"))

	var relTo = target.RelTo(base)
	assertGoPathEqual(t, relTo, FromPath("b/c"))

}

func TestRel_fail(t *testing.T) {

	var base = FromPath("/a")
	var target = FromPath("b/c")

	var rel = base.Rel(target)
	assertGoPathHasErr(t, rel)

	var relTo = target.RelTo(base)
	assertGoPathHasErr(t, relTo)

}
