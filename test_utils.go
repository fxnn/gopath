package gopath

import "testing"

func assertGoPathHasErr(t *testing.T, actual GoPath) {
	if !actual.HasErr() {
		t.Errorf("Expected to have Err(), but was %v", actual)
	}
}

func assertGoPathEqual(t *testing.T, actual GoPath, expected GoPath) {
	if actual.Err() != expected.Err() {
		t.Errorf("expected Err() to be %v, but was %v", expected.Err(), actual.Err())
	}
	if actual.Path() != expected.Path() {
		t.Errorf("expected Path() to be %v, but was %v", expected.Path(), actual.Path())
	}
}

func assertStrEqual(t *testing.T, actual string, expected string) {
	if expected != actual {
		t.Errorf("expected to be equal %s, but was %s", expected, actual)
	}
}
