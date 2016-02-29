package main

import (
	"github.com/nevet/example/errpkg"
	"github.com/nevet/graberror"
)

func main() {
	// a normal error handling paradigm
	err := errpkg.ErrorCustomHandlerDemo()
	if err != nil {
		// you need to handle the error here using a way that's appropriate to YOU,
		// but not to the AUTHOR of the API
		println(err.Error())
	}

	// a graberror handling paradigm
	err := errpkg.ErrorCustomHandlerDemo()
	if err != nil {
		if ge, ok := err.(graberror.GrabError); ok {
			// handle the error using a way that's appropriate to the AUTHOR of the API
			ge.Handle("PkgMain", "FuncMain")
		} else {
			// if the err thrown is not a graberror, handle in a normal way
			println("example", "normal error caught: "+err.Error())
		}
	}
}
