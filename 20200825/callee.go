package sample

import "testing"

func doTest(t *testing.T, canSuccess bool) {
	if !canSuccess {
		t.Error("failed")
	}
}

func doTestUsingHelper(t *testing.T, canSuccess bool) {
	t.Helper()
	if !canSuccess {
		t.Error("failed")
	}
}
