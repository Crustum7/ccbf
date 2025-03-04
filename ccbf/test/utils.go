package test

import "testing"

func ShouldPanic(t *testing.T, f func()) {
	t.Helper()
	defer func() { recover() }()
	f()
	t.Errorf("Should have panicked")
}
