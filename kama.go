// kama: exposes exported information of types, variables, packages, modules, imports etc
//
// It can be used to aid debugging and testing.
//
package kama

import (
	"fmt"
	"reflect"
)

// TODO: If someone passes in, say a struct;
// we should show them its type, methods etc
// but also print it out and its contents
// basically, do what `litter.Dump` would have done

// TODO: clean up

// TODO: fuzz test

// TODO: add documentation for `kama`

// TODO: add a command line api.
//   eg; `kama http.Request` or `kama http`
// have a look at `golang.org/x/tools/cmd/godex`

// Dir returns exported information of types, variables, packages, modules, imports
//
// It is almost similar to Python's builtin dir function
//
// examples:
//
//     import "github.com/komuw/kama"
//
//     kama.Dir("compress/flate")
//     kama.Dir(&http.Request{})
//     kama.Dir("github.com/pkg/errors")
//     kama.Dir(http.Request{})
//
func Dir(i interface{}) string {
	iType := reflect.TypeOf(i)
	if iType == nil {
		res := newVari(i)
		return res.String()
	} else if iType.Kind() == reflect.String {
		i := i.(string)
		res, err := newPak(i)
		if err != nil {
			res := newVari(i)
			return res.String()
		}
		return res.String()
	} else {
		res := newVari(i)
		return res.String()
	}
}

// Dirp prints exported information of types, variables, packages, modules, imports
func Dirp(i interface{}) {
	fmt.Println(Dir(i))
}