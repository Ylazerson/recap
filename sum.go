// B"H

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
Note, this is a command line tool. So `go install` it first.

➜  12 git:(master) ✗ cd 12a-sum

➜  12a-sum git:(master) ✗ go install

➜  12a-sum git:(master) ✗ 12a-sum /home/laz/tmp/ch12-data.txt

Opening /home/laz/tmp/ch12-data.txt
Closing file
Sum: 50.75

➜  12a-sum git:(master) ✗ 12a-sum /home/laz/tmp/ch12-bad-data.txt

Opening /home/laz/tmp/ch12-bad-data.txt
Closing file
2019/09/08 05:26:55 strconv.ParseFloat: parsing "hello": invalid syntax

*/

// -- --------------------------------------
func OpenFile(fileName string) (*os.File, error) {

	fmt.Println("Opening", fileName)

	return os.Open(fileName)
}

// -- --------------------------------------
func CloseFile(file *os.File) {

	fmt.Println("Closing file")

	file.Close()
}

// -- --------------------------------------
// If any errors are encountered along the way, they’ll be returned
// from the GetFloats function, and we’ll store them in the err variable.
func GetFloats(fileName string) ([]float64, error) {

	var numbers []float64

	// -- ----------------------------------
	file, err := OpenFile(fileName)

	if err != nil {
		return nil, err
	}

	// -- ----------------------------------
	defer CloseFile(file)

	// -- ----------------------------------
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		number, err := strconv.ParseFloat(scanner.Text(), 64)

		if err != nil {
			return nil, err
		}

		numbers = append(numbers, number)
	}

	// -- ----------------------------------
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return numbers, nil
}

func main() {

	/*
		Get the name of the file to open from the first command-line argument by accessing os.Args[1].
		Remember, the os.Args[0] element is the name of the program being run.
		The actual program arguments appear in os.Args[1] and later elements.
	*/
	numbers, err := GetFloats(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0

	for _, number := range numbers {
		sum += number

	}

	fmt.Printf("Sum: %0.2f\n", sum)
}
