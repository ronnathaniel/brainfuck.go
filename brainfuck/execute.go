
package brainfuck

import (
	"fmt"
	"strings"
)

func Exec(inputRaw string) {

	t := NewTape()
	Log("THIS IS NEW EXEC")
	for t.windup < uint32(len(inputRaw)) {
		letter := string(inputRaw[t.windup])
		op := getOp(t, letter)
		t.do(op)
	}

	Log("THIS IS NEW")
	fmt.Println(t.stdout)
}

func getOp(t *tape, s string) func() error {
	if ! isValidToken(s) {
		return t.noOp
	}

	var charToOp = map[string]func() error {
		"+": t.increment,
		"-": t.decrement,
		">": t.shiftRight,
		"<": t.shiftLeft,
		"[": t.openLoop,
		"]": t.closeLoop,
		".": t.outputByte,
		",": t.inputByte,
	}

	return charToOp[s]
}

func isValidToken(s string) bool {
	return strings.Contains("+-><,[].,", s)
}

func Log(s ...interface{}) {
	if DEBUG {
		fmt.Println("[BF DEBUG]", s)
	}
}

func Check(err error) {
	if err != nil {
		Log("ERROR - ", err)
		panic(err)
	}
}