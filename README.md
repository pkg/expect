# expect [![Travis-CI](https://travis-ci.org/pkg/expect.svg)](https://travis-ci.org/pkg/expect) [![GoDoc](https://godoc.org/github.com/pkg/expect?status.svg)](http://godoc.org/github.com/pkg/expect) [![Report card](https://goreportcard.com/badge/github.com/pkg/expect)](https://goreportcard.com/report/github.com/pkg/expect) [![Sourcegraph](https://sourcegraph.com/github.com/pkg/expect/-/badge.svg)](https://sourcegraph.com/github.com/pkg/expect?badge)

A simple assertion library that you probably shouldn't use.

## Quickstart

Package `expect` contains various test assertion helpers.
```go
func TestOpenFile(t *testing.T) {
        f, err := os.Open("notfound")
        expect.Nil(err)
        err = f.Close()
        expect.True(err == nil)
}
```
`expect` helpers can be called from any function called from the main testing goroutine.
```bash
% go test
--- FAIL: TestOpenFile (0.00s)
    check_test.go:12: expected: <nil>, got: open notfound: no such file or directory
``` 

Consult the [documentation](https://godoc.org/github.com/pkg/expect) for more information.

## How does this work?

![Magic](https://dave.cheney.net/paste/magic.gif)

## No, seriously, how does this work?

`*testing.T` is recovered dynamically from the call stack. Read the [blog post](https://dave.cheney.net/2019/12/08/dynamically-scoped-variables-in-go).

## License

BSD-2-Clause
