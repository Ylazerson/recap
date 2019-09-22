// B"H

/*
go install github.com/Ylazerson/recap/recap

recap
    COMMAND          : find
	ARG 1 (dir)      : ~/repos/go-workspace/src/github.com/Ylazerson/go-shenanigans/head-first-go/12
	ARG 2 (andOr)    : and
	ARG 3 (keyWords) : fmt err

recap
    COMMAND          : list
	ARG 1 (dir)      : ~/repos/go-workspace/src/github.com/Ylazerson/go-shenanigans/head-first-go/12

*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// -- ------------------------------------------
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

// -- ------------------------------------------
func scanDirectory(path string, filePathsPointer *[]string) {

	// -- --------------------------------------
	// fmt.Println(path)

	files, err := ioutil.ReadDir(path)

	if err != nil {
		panic(err)
	}

	for _, file := range files {

		filePath := filepath.Join(path, file.Name())

		// Note the recursive function call:
		if file.IsDir() {
			scanDirectory(filePath, filePathsPointer)
		} else {
			*filePathsPointer = append(*filePathsPointer, filePath)
		}
	}

}

// -- ------------------------------------------
func main() {

	// -- --------------------------------------
	defer reportPanic()

	// -- --------------------------------------
	// Get the command-line args:
	dir := os.Args[1]
	keyWords := os.Args[2:]

	// -- --------------------------------------
	fmt.Printf("\nDirectory to search in: %#v\n", dir)
	fmt.Printf("\nSearch keywords slice: %#v\n", keyWords)

	// -- --------------------------------------
	keyWordsStr := strings.Join(keyWords, "|")
	fmt.Printf("\nSearch keywords string: %#v\n", keyWordsStr)

	// -- --------------------------------------
	var filePaths []string
	scanDirectory(dir, &filePaths)
	// fmt.Printf("\nAll filePaths: %#v\n", filePaths)

	// -- --------------------------------------
	// add a (?i) at the beginning to make it case insensitive.
	r := regexp.MustCompile("(?i)\\b(" + keyWordsStr + ")\\b")

	for _, filePath := range filePaths {

		// The method takes an integer argument n;
		// if n >= 0, the function returns at most n matches.
		matches := r.FindAllString(filePath, -1)

		if len(matches) > 0 {
			fmt.Println(filePath)
		}
	}

}
