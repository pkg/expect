package expect

import (
	"bufio"
	"bytes"
	"fmt"
	"runtime"
	"testing"
	"unsafe"
)

// getT returns the address of the testing.T passed to testing.tRunner
// which called the function which called getT. If testing.tRunner cannot
// be located in the stack, say if getT is not called from the main test
// goroutine, getT returns nil.
func getT() *testing.T {
	var buf [8192]byte
	n := runtime.Stack(buf[:], false)
	sc := bufio.NewScanner(bytes.NewReader(buf[:n]))
	for sc.Scan() {
		var p uintptr
		n, _ := fmt.Sscanf(sc.Text(), "testing.tRunner(%v", &p)
		if n != 1 {
			continue
		}
		return (*testing.T)(unsafe.Pointer(p))
	}
	return nil
}
