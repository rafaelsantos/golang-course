package tests

import (
	"runtime"
	"testing"
)

func TestDependent(t *testing.T) {
	//Running test in parallel
	t.Parallel()

	if runtime.GOARCH == "amd64" {
		t.Skip("Not works at amd64 architecture")
	}

	t.Fail()
}
