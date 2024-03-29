//go:build unassert_test
// +build unassert_test

package unassert

import (
	"reflect"
	"testing"
)

func TestError(t *testing.T) {
	lastError = errorArgs{}

	Error("test '%s'", "Error")

	assertLastError(t, "test '%s'", "Error")
}

func TestTrue_True(t *testing.T) {
	lastError = errorArgs{}

	True(true)

	assertLastErrorEmpty(t)
}

func TestTrue_False(t *testing.T) {
	lastError = errorArgs{}

	True(false)

	assertLastError(t, "assertion failed: expected true, got false")
}

func TestFalse_False(t *testing.T) {
	lastError = errorArgs{}

	False(false)

	assertLastErrorEmpty(t)
}

func TestFalse_True(t *testing.T) {
	lastError = errorArgs{}

	False(true)

	assertLastError(t, "assertion failed: expected false, got true")
}

func TestNil_Nil(t *testing.T) {
	lastError = errorArgs{}

	Nil(nil)

	assertLastErrorEmpty(t)
}

func TestNil_NotNil(t *testing.T) {
	lastError = errorArgs{}

	v := &struct{}{}
	Nil(v)

	assertLastError(t, "assertion failed: expected nil, got %v", v)
}

func TestNotNil_NotNil(t *testing.T) {
	lastError = errorArgs{}

	NotNil(&struct{}{})

	assertLastErrorEmpty(t)
}

func TestNotNil_Nil(t *testing.T) {
	lastError = errorArgs{}

	NotNil(nil)

	assertLastError(t, "assertion failed: expected non-nil, got nil")
}

func TestSame_Same(t *testing.T) {
	lastError = errorArgs{}

	Same("a", "a")

	assertLastErrorEmpty(t)
}

func TestSame_NotSame(t *testing.T) {
	lastError = errorArgs{}

	Same("a", "b")

	assertLastError(t, "assertion failed: expected same values, got %v != %v", "a", "b")
}

func TestNotSame_Same(t *testing.T) {
	lastError = errorArgs{}

	NotSame("a", "b")

	assertLastErrorEmpty(t)
}

func TestNotSame_NotSame(t *testing.T) {
	lastError = errorArgs{}

	NotSame("a", "a")

	assertLastError(t, "assertion failed: expected two different values, got %v", "a", "a")
}

func TestMatches_True(t *testing.T) {
	lastError = errorArgs{}

	Matches(func(x interface{}) bool { return x == "a" }, "a")

	assertLastErrorEmpty(t)
}

func TestMatches_False(t *testing.T) {
	lastError = errorArgs{}

	Matches(func(x interface{}) bool { return x == "a" }, "b")

	assertLastError(t, "assertion failed: predicate(x) did not return true; x = %v", "b")
}

func TestReturnsTrue_True(t *testing.T) {
	lastError = errorArgs{}

	ReturnsTrue(func() bool { return true })

	assertLastErrorEmpty(t)
}

func TestReturnsTrue_False(t *testing.T) {
	lastError = errorArgs{}

	ReturnsTrue(func() bool { return false })

	assertLastError(t, "assertion failed: evaluator() did not return true")
}

func assertLastError(t *testing.T, format string, v ...interface{}) {
	if format != lastError.format {
		t.Errorf("expected format %q, got %q", format, lastError.format)
	}

	if !reflect.DeepEqual(v, lastError.v) {
		t.Errorf("expected error %q, got %q", format, lastError.format)
	}
}

func assertLastErrorEmpty(t *testing.T) {
	if lastError.format != "" {
		t.Errorf("expected no error, got %q", lastError.format)
	}

	if lastError.v != nil {
		t.Errorf("expected no error args, got %q", lastError.v)
	}
}
