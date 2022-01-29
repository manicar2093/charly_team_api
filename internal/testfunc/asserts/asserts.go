package asserts

import "testing"

func ShouldNotPanic(t *testing.T) {
	if r := recover(); r != nil {
		t.Log(r)
		t.Fatalf("should not panic. recover error: %v", r)
	}
}

func ShouldPanic(t *testing.T, message string) {
	if r := recover(); r == nil {
		t.Fatalf("should panic. Message: %v", message)
	}
}
