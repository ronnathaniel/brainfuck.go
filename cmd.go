
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
	//inputRaw = "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."
	//parse(inputRaw)

	inputRaw = parseArgs()
	brainfuck.Exec(inputRaw)
}

func parseArgs() string {
	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("\tbrainfuck <path/to/file>.bf")
		fmt.Println("\tbrainfuck --exec <bf logic>")
	} else if len(os.Args) == 2 {
		return OpenFile(os.Args[1])
	} else if len(os.Args) == 3 {
		return os.Args[2]
	}

	return ""
}

func OpenFile(path string) string {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	return string(f)
}