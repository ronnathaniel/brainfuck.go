
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ronnathaniel/brainfuck.go/brainfuck"
)

var (
	inputRaw string
)

func main() {

	inputRaw = parseArgs()
	brainfuck.Exec(inputRaw)
}

func parseArgs() string {
	brainfuck.Log("RAW ARGS - ", os.Args)
	var s string

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("\tbrainfuck <path/to/file>.bf")
		fmt.Println("\tbrainfuck --exec <bf logic>")
	} else if len(os.Args) == 2 {
		s = OpenFile(os.Args[1])
	} else if len(os.Args) == 3 {
		s = os.Args[2]
	}

	brainfuck.Log("RAW STRING - ", s)
	return s
}

func OpenFile(path string) string {
	brainfuck.Log("getting file")
	f, err := ioutil.ReadFile(path)
	brainfuck.Check(err)

	return string(f)
}
