# testinglogur

[![GoDoc](https://godoc.org/github.com/muir/logur?status.png)](https://pkg.go.dev/github.com/muir/testinglogur)

Install:

	go get github.com/muir/testinglogur

---

There is a proliferation of Go logging packages.  Most are pretty similar and
just provide ways to manage text logs with levels, colors, and various output
adaptors.  Many allow tagging and inheriting of tags.

There are a few logging packages that allow rich data to be logged.

Choosing a logger is a touch decision for a project.  But, what if your project
is not a app, but rather an open-source library?  Which logger do you use then?
The README for [logur](https://github.com/logur/logur) has an answer. Use the
following:

```go
type MyLogger interface {
	Trace(msg string, fields ...map[string]interface{})
	Debug(msg string, fields ...map[string]interface{})
	Info(msg string, fields ...map[string]interface{})
	Warn(msg string, fields ...map[string]interface{})
	Error(msg string, fields ...map[string]interface{})
}
```

There are adaptors to make a number of the more interesting logging packages
conform to this interface so using this interface for your library still gives
people who use your library the freedom to choose the logger they like best.

So that still leaves one gap: how do you test your library?

This package, `testinglogur` is the answer:

```go
import "github.com/muir/testinglogur"

func TestMyTest(t *testing.T) {
	log := testinglogur.Get(t)
	TestYourLibrary(log)
}
```
