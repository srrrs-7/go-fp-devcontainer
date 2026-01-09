package task

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTestFunc(t *testing.T) {
	expected := "This is a test function in func.go"
	result := TestFunc()
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Errorf("TestFunc() mismatch (-want +got):\n%s", diff)
	}
}
