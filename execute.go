
package main

import (
	"fmt"
	"strings"
)

func Exec(inputRaw string) {

	t := NewTape()
	charToOp := map[string]func() error {
		"+": t.increment,
		"-": t.decrement,
		">": t.shiftRight,
		"<": t.shiftLeft,
		"[": t.openLoop,
		"]": t.closeLoop,
		".": t.outputByte,
		",": t.inputByte,
	}

	for t.windup < uint32(len(inputRaw))  {
		letter := string(inputRaw[t.windup])

		if ! isValidToken(letter) {
			continue
		}

		op := charToOp[letter]
		t.do(op)
	}

	fmt.Println(t.stdout)

}

func isValidToken(s string) bool {
	return strings.Contains("+-><,[].,", s)
}

func log(s ...interface{}) {
	if DEBUG {
		fmt.Println("[BF DEBUG]", s)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}