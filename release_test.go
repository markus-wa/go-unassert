// +build !unassert_stderr,!unassert_panic,!unassert_test

package unassert

import (
	"testing"
)

func TestError(t *testing.T) {
	Error("test '%s'", "Error")
}

func TestTrue_False(t *testing.T) {
	True(false)
}

func TestTruef_NOP(t *testing.T) {
	Truef(false, "Truef")
}

func TestFalse_NOP(t *testing.T) {
	False(false)
}

func TestFalsef_NOP(t *testing.T) {
	Falsef(false, "Falsef")
}

func TestNil_NOP(t *testing.T) {
	Nil(&struct{}{})
}

func TestNilf_NOP(t *testing.T) {
	Nilf(&struct{}{}, "Nilf")
}

func TestNotNil_NOP(t *testing.T) {
	NotNil(nil)
}

func TestNotNilf_NOP(t *testing.T) {
	NotNilf(nil, "NotNilf")
}

func TestSame_NOP(t *testing.T) {
	Same("a", "b")
}

func TestSamef_NOP(t *testing.T) {
	Samef("a", "b", "Samef")
}

func TestNotSame_NOP(t *testing.T) {
	NotSame("a", "a")
}

func TestNotSamef_NOP(t *testing.T) {
	NotSamef("a", "a", "NotSamef")
}

func TestMatches_NOP(t *testing.T) {
	Matches(nil, nil)
}

func TestMatchesf_NOP(t *testing.T) {
	Matchesf(nil, nil, "Matchesf")
}

func TestReturnsTrue_NOP(t *testing.T) {
	ReturnsTrue(nil)
}

func TestReturnsTruef_NOP(t *testing.T) {
	ReturnsTruef(nil, "ReturnsTruef")
}
