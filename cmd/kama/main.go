package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/komuw/kama"
)

// TODO: clean up

// TODO: fuzz test

// TODO: add documentation for `kama`

// TODO: add a command line api.
//   eg; `kama http.Request` or `kama http`
// have a look at `golang.org/x/tools/cmd/godex`

func usage() {
	fmt.Fprintf(os.Stderr,
		`kama prints exported information of types, variables, packages, modules, imports etc.
It also pretty prints data structures.

examples:
	kama github.com/pkg/errors
	    prints exported information of the package(github.com/pkg/errors)
	kama fmt.Printf
	    prints exported information of the package object(fmt.Printf)
	kama http.Response
	    prints exported information of the package object(http.Response)`)
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 1 {
		// TODO: suppport more args.
		// see: https://github.com/golang/tools/blob/gopls/v0.5.1/cmd/godex/godex.go#L58
		fmt.Fprintln(os.Stderr, "error: number of cli arguments should be one")
		os.Exit(2)
	}

	args := flag.Args()
	kama.Dirp(args[0])
}
