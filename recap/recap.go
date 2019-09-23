// B''H

package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"text/template"
)

// -- -----------------------------------------
// check calls log.Fatal on any non-nil error.
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// -- -----------------------------------------
type DirStruct struct {
	Dirs []string
}

// -- -----------------------------------------
func viewHandler(writer http.ResponseWriter, request *http.Request) {

	dirs := getStrings("dirs.txt")

	html, err := template.ParseFiles("view.html")
	check(err)

	dirStruct := DirStruct{
		Dirs: dirs,
	}

	err = html.Execute(writer, dirStruct)
	check(err)
}

// -- -----------------------------------------
// getStrings returns a slice of strings read from fileName, one
// string per line.
func getStrings(fileName string) []string {

	var lines []string

	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())

	return lines
}

// -- -----------------------------------------
func main() {
	http.HandleFunc("/recap", viewHandler)

	err := http.ListenAndServe("localhost:8080", nil)

	log.Fatal(err)
}
