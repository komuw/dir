## kama          

[![ci](https://github.com/komuw/kama/workflows/kama%20ci/badge.svg)](https://github.com/komuw/kama/actions)
[![codecov](https://codecov.io/gh/komuw/kama/branch/master/graph/badge.svg)](https://codecov.io/gh/komuw/kama)


`kama` prints exported information of types, variables, packages, modules, imports etc     
It also pretty prints data structures.    
It can be used to aid debugging and testing.        

It's name is derived from Kenyan hip hop artiste, `Kama`(One third of the hiphop group `Kalamashaka`).                               

Comprehensive documetion is available -> [Documentation](https://pkg.go.dev/github.com/komuw/kama)


## Installation

```shell
go get github.com/komuw/kama # library
go get github.com/komuw/kama/cmd/kama # cli app
```           


## Usage

#### 1. As a library

```go
import "github.com/komuw/kama"

kama.Dirp("compress/flate")
kama.Dirp(&http.Request{})
kama.Dirp("github.com/pkg/errors")
kama.Dirp(http.Request{})
```

This;   
```go
import "github.com/komuw/kama"

kama.Dirp("github.com/pkg/errors")
```
will output;   
```bash
[
NAME: github.com/pkg/errors
CONSTANTS: []
VARIABLES: []
FUNCTIONS: [
	As(err error, target interface{}) bool
	Cause(err error) error
	Errorf(format string, args ...interface{}) error
	Is(err error, target error) bool
	New(message string) error
	Unwrap(err error) error
	WithMessage(err error, message string) error
	WithMessagef(err error, format string, args ...interface{}) error
	WithStack(err error) error
	Wrap(err error, message string) error
	Wrapf(err error, format string, args ...interface{}) error
	]
TYPES: [
	Frame uintptr
		(Frame) Format(s fmt.State, verb rune)
		(Frame) MarshalText() ([]byte, error)
	StackTrace []Frame
		(StackTrace) Format(s fmt.State, verb rune)]
]
```   
   
  
whereas this;   
```go
import "github.com/komuw/kama"

h := http.Header{}
h.Add("content-type", "text")
kama.Dirp(h)
```
will output;  
```bash
[
NAME: net/http.Header
KIND: map
SIGNATURE: [http.Header]
FIELDS: []
METHODS: [
	Add func(http.Header, string, string)
	Clone func(http.Header) http.Header
	Del func(http.Header, string)
	Get func(http.Header, string) string
	Set func(http.Header, string, string)
	Values func(http.Header, string) []string
	Write func(http.Header, io.Writer) error
	WriteSubset func(http.Header, io.Writer, map[string]bool) error
	]
SNIPPET: Header{
  "Content-Type": []string{
    "text",
  },
}
]
```

#### 2. As a cli app    
`kama` also has a commandline app, which you can install as;
```shell
go get github.com/komuw/kama/cmd/kama
```
and use as;
```shell
kama --help
kama github.com/pkg/errors
```


## Inspiration
1. Python's [`dir`](https://docs.python.org/3/library/functions.html#dir) builtin.    
2. [`godex`](https://pkg.go.dev/golang.org/x/tools/cmd/godex).   
3. [`sanity-io/litter`](https://github.com/sanity-io/litter).
