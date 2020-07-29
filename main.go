package main

import (
	"fmt"
	"net/http"

	"reflect"
)

// TODO: If someone passes in, say a struct;
// we should show them its type, methods etc
// but also print it out and its contents
// basically, do what `litter.Dump` would have done

// TODO: clean up

// TODO: add documentation for `dir`

// TODO: add a command line api.
//   eg; `dir http.Request` or `dir http`
// have a look at `golang.org/x/tools/cmd/godex`

// TODO: this will stutter; `dir.dir(23)`
// maybe it is okay??
// TODO: surface all info for both the type and its pointer.
// currently `dir(&http.Client{})` & `dir(http.Client{})` produces different output; they should NOT
func dir(i interface{}) {
	var res interface{}
	var err error

	if reflect.TypeOf(i).Kind() == reflect.String {
		i := i.(string)
		res, err = newPaki(i)
		if err != nil {
			res = newVari(i)
		}
	} else {
		res = newVari(i)
	}

	fmt.Println(res)

}

func main() {
	defer panicHandler()

	dir("archive/tar")
	dir("compress/flate")
	dir(&http.Request{})
	dir(http.Request{})
	dir("github.com/pkg/errors")
}
