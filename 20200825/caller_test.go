package sample

import "testing"

func TestSample(t *testing.T) {
	doTest(t, true)
	doTest(t, true)
	doTest(t, true)
	doTest(t, true)
	doTest(t, false) // <- Test result does't show this line as below because doTestUsingHelper function does't use t.Helper()
	// --- FAIL: TestSample (0.00s)
	// ~~/github.com/blck-snwmn/sandbox4go/20200825/callee.go:7: failed
}

func TestSampleUsingHelper(t *testing.T) {
	doTestUsingHelper(t, true)
	doTestUsingHelper(t, true)
	doTestUsingHelper(t, true)
	doTestUsingHelper(t, true)
	doTestUsingHelper(t, false) // <- Test result shows this line as below because doTestUsingHelper function use t.Helper()
	// print
	// --- FAIL: TestSampleUsingHelper (0.00s)
	// ~~~/github.com/blck-snwmn/sandbox4go/20200825/caller_test.go:20: failed
}
