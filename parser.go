
package main

import (
	"fmt"
	"strings"
)

func parse(inputRaw string) {

	t := NewTape()
	var charToOp = map[string]func() error{
		"+": t.increment,
		"-": t.decrement,
		">": t.shiftRight,
		"<": t.shiftLeft,
		"[": t.openLoop,
		"]": t.closeLoop,
		".": t.outputByte,
		",": t.inputByte,
	}
	//fmt.Println(inputRaw)

	for t.windup < uint32(len(inputRaw)) {
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