// B"H

package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// -- -------------------------------------
/*
When reportPanic is called (see main()), we don’t know
whether the program is actually panicking or not.

The deferred call to reportPanic will be made regardless
of whether scanDirectory calls panic or not.

So the first thing we do is test whether the panic value
returned from recover is nil.

*/
func reportPanic() {

	p := recover()

	if p == nil {
		return
	}

	err, ok := p.(error)

	if ok {
		fmt.Println(err)
	} else {
		panic(p)
	}
}

// -- -------------------------------------
func scanDirectory(path string) {

	fmt.Println(path)

	files, err := ioutil.ReadDir(path)

	if err != nil {
		panic(err)
	}

	for _, file := range files {

		filePath := filepath.Join(path, file.Name())

		// Note the recursive function call:
		if file.IsDir() {
			scanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}

// -- -------------------------------------
func main() {

	defer reportPanic()

	// Uncomment next line to test panic that can’t be converted to an `error` type.
	//panic("some other issue")

	scanDirectory("/home/laz/repos/go-workspace/src/github.com/Ylazerson/go-shenanigans/head-first-go/12")
}
