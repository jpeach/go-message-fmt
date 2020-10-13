# go-message-fmt

Go linter for log message formatting.

Current checks:

* that logrus log messages are not capitalized
* that kingpin command help is capitalized

## Installation

```
$ go install .
```

## Running

```
$ go-message-fmt ./...
/Users/jpeach/upstream/contour/internal/k8s/status.go:123:12: message starts with uppercase: "Update was a no-op"
...
```
