// package expect contains various test assertion helpers.
//
// The functions this package recover the *testing.T value passed
// to the Test from the call stack. This implies that expect
// functions should not be called from a goroutine launched from
// a test. However, this restriction applies to calling t.Fatal/FailNow
// directly, so here we are.
package expect

// Nil fails the test if v is not nil.
func Nil(v interface{}) {
	if v != nil {
		t := getT()
		t.Helper()
		t.Fatalf("expected: %v, got: %v", nil, v)
	}
}

// True fails the test if v is not true.
func True(v bool) {
	if !v {
		t := getT()
		t.Helper()
		t.Fatalf("expected: %v, got: %v", true, false)
	}
}
